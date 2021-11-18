package cmd

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type registry struct{}

type params struct {
	insecure bool
	auth     authentication.Client
}

type gRPCClient struct{}

func newRegistry(params *params) (*registry, error) {
	return &registry{}, nil
}

func newGRPCClient(params *params) (*gRPCClient, error) {
	_, err := newGRPCOptions(params)
	if err != nil {
		return nil, err
	}

	return &gRPCClient{}, nil
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
