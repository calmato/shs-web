package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	"github.com/calmato/shs-web/api/internal/classroom/api"
	db "github.com/calmato/shs-web/api/internal/classroom/database"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type registry struct {
	classroomServer classroom.ClassroomServiceServer
}

type params struct {
	insecure       bool
	logger         *zap.Logger
	db             *database.Client
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

	dbParams := &db.Params{
		Database: params.db,
	}
	apiParams := &api.Params{
		Logger:      params.logger,
		WaitGroup:   &sync.WaitGroup{},
		Database:    db.NewDatabase(dbParams),
		UserService: cli.user,
	}
	return &registry{
		classroomServer: api.NewClassroomService(apiParams),
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
