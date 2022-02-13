package notifier

import (
	"context"
	"errors"
	"net/url"
	"os"
	"testing"
	"time"

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
	url, err := url.Parse("http://example.com")
	require.NoError(t, err)
	params := &Params{
		TeacherWebURL: url,
		StudentWebURL: url,
	}
	assert.NotNil(t, NewNotifier(params))
}

func TestNotifier_Run(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mocks := newMocks(ctrl)
	notifier := newNotifier(mocks)
	msgCh := make(chan *pubsub.Message, 1)
	go func() {
		assert.NoError(t, notifier.Run(ctx, msgCh, 1, 1))
	}()
	time.Sleep(3 * time.Second)
	cancel()
}

func TestNotifier_Dispatch(t *testing.T) {
	t.Parallel()

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
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *messenger.NotifierRequest
		expect bool
		isErr  bool
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
			req: &messenger.NotifierRequest{
				Key:        uuid.Base58Encode(uuid.New()),
				TeacherIds: []string{"teacherid"},
				StudentIds: []string{"studentid"},
				Email: &messenger.EmailConfig{
					EmailId:       "notifier-test",
					Substitutions: map[string]string{"key": "value"},
				},
			},
			expect: false,
			isErr:  false,
		},
		{
			name:  "success to email config is empty",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {},
			req: &messenger.NotifierRequest{
				Key:        uuid.Base58Encode(uuid.New()),
				TeacherIds: []string{"teacherid"},
				StudentIds: []string{"studentid"},
			},
			expect: false,
			isErr:  false,
		},
		{
			name: "failed to get teachers and students",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(nil, errors.New("some error"))
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(nil, errors.New("some error"))
			},
			req: &messenger.NotifierRequest{
				Key:        uuid.Base58Encode(uuid.New()),
				TeacherIds: []string{"teacherid"},
				StudentIds: []string{"studentid"},
				Email: &messenger.EmailConfig{
					EmailId:       "notifier-test",
					Substitutions: map[string]string{"key": "value"},
				},
			},
			expect: true,
			isErr:  true,
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
			req: &messenger.NotifierRequest{
				Key:        uuid.Base58Encode(uuid.New()),
				TeacherIds: []string{"teacherid"},
				StudentIds: []string{"studentid"},
				Email: &messenger.EmailConfig{
					EmailId:       "notifier-test",
					Substitutions: map[string]string{"key": "value"},
				},
			},
			expect: false,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mocks := newMocks(ctrl)

			tt.setup(ctx, t, mocks)

			notifier := newNotifier(mocks)

			actual, err := notifier.dispatch(ctx, tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func setEnv() {
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", "127.0.0.1:8085")
	}
}
