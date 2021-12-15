package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	v1 "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/handler"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type registry struct {
	v1 v1.APIV1Handler
}

type params struct {
	insecure       bool
	logger         *zap.Logger
	auth           authentication.Client
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

	v1Params := &v1.Params{
		Auth:        params.auth,
		UserService: cli.user,
		Logger:      params.logger,
		WaitGroup:   &sync.WaitGroup{},
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
