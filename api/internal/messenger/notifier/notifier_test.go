package notifier

import (
	"context"
	"errors"
	"net/url"
	"os"
	"testing"
	"time"

	gpubsub "cloud.google.com/go/pubsub"
	"github.com/calmato/shs-web/api/internal/messenger/mailer"
	mock_mailer "github.com/calmato/shs-web/api/mock/messenger/mailer"
	mock_pubsub "github.com/calmato/shs-web/api/mock/pkg/pubsub"
	mock_user "github.com/calmato/shs-web/api/mock/proto/user"
	"github.com/calmato/shs-web/api/pkg/pubsub"
	"github.com/calmato/shs-web/api/pkg/uuid"
	"github.com/calmato/shs-web/api/proto/messenger"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const topicID = "notifier-test"

type mocks struct {
	mailer *mock_mailer.MockClient
	puller *mock_pubsub.MockPuller
	user   *mock_user.MockUserServiceClient
}

type testOptions struct{}

type testOption func(*testOptions)

func newMocks(ctrl *gomock.Controller) *mocks {
	return &mocks{
		mailer: mock_mailer.NewMockClient(ctrl),
		puller: mock_pubsub.NewMockPuller(ctrl),
		user:   mock_user.NewMockUserServiceClient(ctrl),
	}
}

func newNotifier(mocks *mocks, opts ...testOption) *Notifier {
	dopts := &testOptions{}
	for i := range opts {
		opts[i](dopts)
	}
	webURL, _ := url.Parse("http://example.com")
	return &Notifier{
		logger: zap.NewNop(),
		mailer: mocks.mailer,
		puller: mocks.puller,
		user:   mocks.user,
		teacherWebURL: func() *url.URL {
			url := *webURL
			return &url
		},
		studentWebURL: func() *url.URL {
			url := *webURL
			return &url
		},
	}
}

func TestNotifier(t *testing.T) {
	t.Parallel()
	assert.NotNil(t, NewNotifier(&Params{}))
}

func TestNotifier_Run(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := pubsub.NewClient(ctx, "test-project")
	require.NoError(t, err)
	topic, err := client.Client.CreateTopic(ctx, topicID)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		t.Fatal(err)
	}
	if status.Code(err) == codes.AlreadyExists {
		topic = client.Client.Topic(topicID)
	}
	conf := gpubsub.SubscriptionConfig{Topic: topic}
	_, err = client.Client.CreateSubscription(ctx, topicID, conf)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		t.Fatal(err)
	}

	publisher := pubsub.NewPublisher(client, topicID, pubsub.WithPublisherMaxRetries(1))
	puller := pubsub.NewPuller(client, topicID, pubsub.WithPullerTimeout(3*time.Second))

	teachers := []*user.Teacher{{
		Id:        "teacherid",
		LastName:  "テスト",
		FirstName: "講師",
		Mail:      "test-teacher@calmato.jp",
	}}
	students := []*user.Student{{
		Id:        "studentid",
		LastName:  "テスト",
		FirstName: "生徒",
		Mail:      "test-student@calmato.jp",
	}}

	tests := []struct {
		name            string
		setup           func(ctx context.Context, t *testing.T, mocks *mocks)
		after           func(ctx context.Context, t *testing.T)
		concurrency     int
		mailboxCapacity int
		isErr           bool
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				teachersOut := &user.MultiGetTeachersResponse{Teachers: teachers}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				studentsOut := &user.MultiGetStudentsResponse{Students: students}
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.mailer.EXPECT().MultiSendFromInfo(gomock.Any(), "notifier-test", gomock.Any()).Return(nil)
			},
			after: func(ctx context.Context, t *testing.T) {
				pb := &messenger.NotifierRequest{
					Key:        uuid.Base58Encode(uuid.New()),
					TeacherIds: []string{"teacherid"},
					StudentIds: []string{"studentid"},
					Email: &messenger.EmailConfig{
						EmailId:       "notifier-test",
						Substitutions: map[string]string{"key": "value"},
					},
				}
				data, err := proto.Marshal(pb)
				require.NoError(t, err)
				_, err = publisher.Publish(ctx, &pubsub.Message{Data: data})
				require.NoError(t, err)
			},
			concurrency:     2,
			mailboxCapacity: 1,
			isErr:           false,
		},
		{
			name: "success after failed to get teachers and students",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				teachersOut := &user.MultiGetTeachersResponse{Teachers: teachers}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				studentsOut := &user.MultiGetStudentsResponse{Students: students}
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(nil, errors.New("some error"))
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(nil, errors.New("some error"))
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.mailer.EXPECT().MultiSendFromInfo(gomock.Any(), "notifier-test", gomock.Any()).Return(nil)
			},
			after: func(ctx context.Context, t *testing.T) {
				pb := &messenger.NotifierRequest{
					Key:        uuid.Base58Encode(uuid.New()),
					TeacherIds: []string{"teacherid"},
					StudentIds: []string{"studentid"},
					Email: &messenger.EmailConfig{
						EmailId:       "notifier-test",
						Substitutions: map[string]string{"key": "value"},
					},
				}
				data, err := proto.Marshal(pb)
				require.NoError(t, err)
				_, err = publisher.Publish(ctx, &pubsub.Message{Data: data})
				require.NoError(t, err)
			},
			concurrency:     2,
			mailboxCapacity: 1,
			isErr:           false,
		},
		{
			name: "success after failed to send email",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				teachersOut := &user.MultiGetTeachersResponse{Teachers: teachers}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				studentsOut := &user.MultiGetStudentsResponse{Students: students}
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Times(2).Return(teachersOut, nil)
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Times(2).Return(studentsOut, nil)
				mocks.mailer.EXPECT().
					MultiSendFromInfo(gomock.Any(), "notifier-test", gomock.Any()).
					Return(mailer.ErrUnavailable)
				mocks.mailer.EXPECT().
					MultiSendFromInfo(gomock.Any(), "notifier-test", gomock.Any()).
					Return(nil)
			},
			after: func(ctx context.Context, t *testing.T) {
				pb := &messenger.NotifierRequest{
					Key:        uuid.Base58Encode(uuid.New()),
					TeacherIds: []string{"teacherid"},
					StudentIds: []string{"studentid"},
					Email: &messenger.EmailConfig{
						EmailId:       "notifier-test",
						Substitutions: map[string]string{"key": "value"},
					},
				}
				data, err := proto.Marshal(pb)
				require.NoError(t, err)
				_, err = publisher.Publish(ctx, &pubsub.Message{Data: data})
				require.NoError(t, err)
			},
			concurrency:     2,
			mailboxCapacity: 1,
			isErr:           false,
		},
		{
			name:  "success to message empty",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {},
			after: func(ctx context.Context, t *testing.T) {
				pb := &messenger.NotifierRequest{
					Key:   uuid.Base58Encode(uuid.New()),
					Email: nil,
				}
				data, err := proto.Marshal(pb)
				require.NoError(t, err)
				_, err = publisher.Publish(ctx, &pubsub.Message{Data: data})
				require.NoError(t, err)
			},
			concurrency:     1,
			mailboxCapacity: 1,
			isErr:           false,
		},
		{
			name:  "failed to invalid request",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {},
			after: func(ctx context.Context, t *testing.T) {
				_, err = publisher.Publish(ctx, &pubsub.Message{Data: []byte("hello")})
				require.NoError(t, err)
			},
			concurrency:     1,
			mailboxCapacity: 1,
			isErr:           false,
		},
		{
			name: "failed to send mail",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				teachersOut := &user.MultiGetTeachersResponse{Teachers: teachers}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				studentsOut := &user.MultiGetStudentsResponse{Students: students}
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.mailer.EXPECT().
					MultiSendFromInfo(gomock.Any(), "notifier-test", gomock.Any()).
					Return(mailer.ErrNotFound)
			},
			after: func(ctx context.Context, t *testing.T) {
				pb := &messenger.NotifierRequest{
					Key:        uuid.Base58Encode(uuid.New()),
					TeacherIds: []string{"teacherid"},
					StudentIds: []string{"studentid"},
					Email: &messenger.EmailConfig{
						EmailId:       "notifier-test",
						Substitutions: map[string]string{"key": "value"},
					},
				}
				data, err := proto.Marshal(pb)
				require.NoError(t, err)
				_, err = publisher.Publish(ctx, &pubsub.Message{Data: data})
				require.NoError(t, err)
			},
			concurrency:     1,
			mailboxCapacity: 1,
			isErr:           false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mocks := newMocks(ctrl)

			tt.setup(ctx, t, mocks)

			notifier := newNotifier(mocks)
			notifier.puller = puller

			eg := errgroup.Group{}
			eg.Go(func() error {
				ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
				defer cancel()
				return notifier.Run(ctx, tt.concurrency, tt.mailboxCapacity)
			})

			tt.after(ctx, t)
			assert.Equal(t, tt.isErr, eg.Wait() != nil, err)
		})
	}
}

func setEnv() {
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", "127.0.0.1:8085")
	}
}
