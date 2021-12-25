package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	v1 "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/handler"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type registry struct {
	v1 v1.APIV1Handler
}

type params struct {
	insecure            bool
	logger              *zap.Logger
	auth                authentication.Client
	classroomServiceURL string
	lessonServiceURL    string
	userServiceURL      string
}

type gRPCClient struct {
	classroom classroom.ClassroomServiceClient
	lesson    lesson.LessonServiceClient
	user      user.UserServiceClient
}

func newRegistry(params *params) (*registry, error) {
	cli, err := newGRPCClient(params)
	if err != nil {
		return nil, err
	}

	v1Params := &v1.Params{
		Auth:             params.auth,
		ClassroomService: cli.classroom,
		LessonService:    cli.lesson,
		UserService:      cli.user,
		Logger:           params.logger,
		WaitGroup:        &sync.WaitGroup{},
	}
	v1Handler := v1.NewAPIV1Handler(v1Params)

	return &registry{
		v1: v1Handler,
	}, nil
}

func newGRPCClient(params *params) (*gRPCClient, error) {
	opts, err := newGRPCOptions(params)
	if err != nil {
		return nil, err
	}

	classroomConn, err := grpc.Dial(params.classroomServiceURL, opts...)
	if err != nil {
		return nil, err
	}
	lessonConn, err := grpc.Dial(params.lessonServiceURL, opts...)
	if err != nil {
		return nil, err
	}
	userConn, err := grpc.Dial(params.userServiceURL, opts...)
	if err != nil {
		return nil, err
	}

	return &gRPCClient{
		classroom: classroom.NewClassroomServiceClient(classroomConn),
		lesson:    lesson.NewLessonServiceClient(lessonConn),
		user:      user.NewUserServiceClient(userConn),
	}, nil
}

func newGRPCOptions(params *params) ([]grpc.DialOption, error) {
	var opts []grpc.DialOption
	if params.insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}
	return opts, nil
}
