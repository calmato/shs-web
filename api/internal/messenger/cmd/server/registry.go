package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	"github.com/calmato/shs-web/api/internal/messenger/api"
	"github.com/calmato/shs-web/api/pkg/pubsub"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/messenger"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type registry struct {
	messengerServer messenger.MessengerServiceServer
}

type params struct {
	insecure         bool
	logger           *zap.Logger
	publisher        pubsub.Publisher
	userServiceURL   string
	lessonServiceURL string
}

type gRPCClient struct {
	user   user.UserServiceClient
	lesson lesson.LessonServiceClient
}

func newRegistry(params *params) (*registry, error) {
	cli, err := newGRPCClient(params)
	if err != nil {
		return nil, err
	}

	apiParams := &api.Params{
		Logger:        params.logger,
		WaitGroup:     &sync.WaitGroup{},
		Publisher:     params.publisher,
		UserService:   cli.user,
		LessonService: cli.lesson,
	}
	return &registry{
		messengerServer: api.NewMessengerService(apiParams),
	}, nil
}

func newGRPCClient(params *params) (*gRPCClient, error) {
	opts, err := newGRPCOptions(params)
	if err != nil {
		return nil, err
	}

	userConn, err := grpc.Dial(params.userServiceURL, opts...)
	if err != nil {
		return nil, err
	}
	lessonConn, err := grpc.Dial(params.lessonServiceURL, opts...)
	if err != nil {
		return nil, err
	}

	return &gRPCClient{
		user:   user.NewUserServiceClient(userConn),
		lesson: lesson.NewLessonServiceClient(lessonConn),
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
