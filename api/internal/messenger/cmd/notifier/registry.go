package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"net/url"

	"github.com/calmato/shs-web/api/internal/messenger/mailer"
	"github.com/calmato/shs-web/api/internal/messenger/notifier"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type registry struct {
	notifier *notifier.Notifier
}

type params struct {
	insecure       bool
	logger         *zap.Logger
	teacherWebURL  *url.URL
	studentWebURL  *url.URL
	mailer         mailer.Client
	userServiceURL string
}

type gRPCClient struct {
	user user.UserServiceClient
}

func newRegistry(params *params) (*registry, error) {
	cli, err := newGRPCClient(params)
	if err != nil {
		return nil, err
	}

	notifierParams := &notifier.Params{
		Logger:        params.logger,
		TeacherWebURL: params.teacherWebURL,
		StudentWebURL: params.studentWebURL,
		Mailer:        params.mailer,
		UserService:   cli.user,
	}
	return &registry{
		notifier: notifier.NewNotifier(notifierParams),
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

	return &gRPCClient{
		user: user.NewUserServiceClient(userConn),
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
