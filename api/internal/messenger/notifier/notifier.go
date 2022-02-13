package notifier

import (
	"context"
	"errors"
	"net/url"
	"sync"

	"github.com/calmato/shs-web/api/internal/messenger/entity"
	"github.com/calmato/shs-web/api/internal/messenger/mailer"
	"github.com/calmato/shs-web/api/pkg/pubsub"
	"github.com/calmato/shs-web/api/proto/messenger"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

type Params struct {
	Logger        *zap.Logger
	TeacherWebURL *url.URL
	StudentWebURL *url.URL
	Mailer        mailer.Client
	Puller        pubsub.Puller
	UserService   user.UserServiceClient
}

type Notifier struct {
	logger        *zap.Logger
	mailer        mailer.Client
	puller        pubsub.Puller
	user          user.UserServiceClient
	teacherWebURL func() *url.URL
	studentWebURL func() *url.URL
}

func NewNotifier(params *Params) *Notifier {
	return &Notifier{
		logger: params.Logger,
		mailer: params.Mailer,
		puller: params.Puller,
		user:   params.UserService,
		teacherWebURL: func() *url.URL {
			url := *params.TeacherWebURL // copy
			return &url
		},
		studentWebURL: func() *url.URL {
			url := *params.StudentWebURL // copy
			return &url
		},
	}
}

func (n *Notifier) Run(
	ctx context.Context, msgCh <-chan *pubsub.Message, concurrency int, mailboxCapacity int,
) error {
	eg, ectx := errgroup.WithContext(ctx)
	// 指定された並列実行数分goroutineを実行
	for i := 0; i < concurrency; i++ {
		eg.Go(func() error {
			return n.run(ectx, msgCh)
		})
	}
	n.logger.Info("Run workers", zap.Int("concurrency", concurrency))
	defer n.logger.Info("Stopped workers")
	return eg.Wait()
}

func (n *Notifier) run(ctx context.Context, msgCh <-chan *pubsub.Message) error {
	for {
		select {
		case <-ctx.Done():
			n.logger.Info("Context is done")
			return nil
		case msg := <-msgCh:
			var req messenger.NotifierRequest
			// リクエスト値の生成に失敗したものについてはリトライせずにACKを返す
			if err := proto.Unmarshal(msg.Data, &req); err != nil {
				n.logger.Error("Failed to unmarshal proto", zap.Error(err))
				msg.Ack()
				continue
			}
			retryable, err := n.dispatch(ctx, &req)
			if err != nil {
				n.logger.Error("Failed to dispatch error", zap.Error(err), zap.String("request.key", req.Key))
			}
			if retryable {
				msg.Nack()
				continue
			}
			msg.Ack()
		}
	}
}

func (n *Notifier) dispatch(ctx context.Context, req *messenger.NotifierRequest) (bool, error) {
	n.logger.Debug("Dispatch Message", zap.Any("msg", req))
	var emailRetryable bool
	var err error
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// メール送信
	go func() {
		defer wg.Done()
		emailRetryable, err = n.sendEmail(ctx, req.TeacherIds, req.StudentIds, req.Email)
		if err == nil {
			return
		}
		if emailRetryable {
			n.logger.Warn("Retry to send email", zap.Error(err), zap.Any("request", req))
			return
		}
		n.logger.Error("Failed to send email", zap.Error(err), zap.Any("request", req))
	}()
	wg.Wait()
	retryable := emailRetryable
	return retryable, err
}

func (n *Notifier) sendEmail(
	ctx context.Context, teacherIDs []string, studentIDs []string, msg *messenger.EmailConfig,
) (bool, error) {
	if msg == nil {
		return false, nil
	}

	eg, ectx := errgroup.WithContext(ctx)
	var teachers entity.Teachers
	eg.Go(func() error {
		in := &user.MultiGetTeachersRequest{Ids: teacherIDs}
		out, err := n.user.MultiGetTeachers(ectx, in)
		if err != nil {
			return err
		}
		teachers = entity.NewTeachers(out.Teachers)
		return nil
	})
	var students entity.Students
	eg.Go(func() error {
		in := &user.MultiGetStudentsRequest{Ids: studentIDs}
		out, err := n.user.MultiGetStudents(ectx, in)
		if err != nil {
			return err
		}
		students = entity.NewStudents(out.Students)
		return nil
	})
	if err := eg.Wait(); err != nil {
		return true, err
	}

	ps := make([]*mailer.Personalization, 0, len(teachers)+len(students))
	ps = append(ps, n.newPersonalizationsForTeacher(teachers, msg)...)
	ps = append(ps, n.newPersonalizationsForStudent(students, msg)...)

	if err := n.mailer.MultiSendFromInfo(ctx, msg.EmailId, ps); err != nil {
		retryable := errors.Is(err, mailer.ErrTimeout) ||
			errors.Is(err, mailer.ErrUnavailable) ||
			errors.Is(err, mailer.ErrInternal)
		return retryable, err
	}
	return false, nil
}

func (n *Notifier) newPersonalizationsForTeacher(
	ts entity.Teachers, msg *messenger.EmailConfig,
) []*mailer.Personalization {
	maker := mailer.NewTeacherURLMaker(n.teacherWebURL())
	webURL := maker.SignIn()

	ps := make([]*mailer.Personalization, len(ts))
	for i := range ts {
		builder := mailer.NewTemplateDataBuilder().
			Data(msg.Substitutions).
			TeacherID(ts[i].Id).
			Name(ts[i].Fullname()).
			WebURL(webURL)
		p := &mailer.Personalization{
			Name:          ts[i].Fullname(),
			Address:       ts[i].Mail,
			Type:          mailer.AddressTypeTo,
			Substitutions: mailer.NewSubstitutions(builder.Build()),
		}
		ps[i] = p
	}
	return ps
}

func (n *Notifier) newPersonalizationsForStudent(
	ss entity.Students, msg *messenger.EmailConfig,
) []*mailer.Personalization {
	maker := mailer.NewStudentURLMaker(n.studentWebURL())
	webURL := maker.SignIn()

	ps := make([]*mailer.Personalization, len(ss))
	for i := range ss {
		builder := mailer.NewTemplateDataBuilder().
			Data(msg.Substitutions).
			StudentID(ss[i].Id).
			Name(ss[i].Fullname()).
			WebURL(webURL)
		p := &mailer.Personalization{
			Name:          ss[i].Fullname(),
			Address:       ss[i].Mail,
			Type:          mailer.AddressTypeTo,
			Substitutions: mailer.NewSubstitutions(builder.Build()),
		}
		ps[i] = p
	}
	return ps
}
