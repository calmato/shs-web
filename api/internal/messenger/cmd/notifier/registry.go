package cmd

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/calmato/shs-web/api/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type params struct {
	insecure       bool
	userServiceURL string
}

type gRPCClient struct {
	user user.UserServiceClient
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
