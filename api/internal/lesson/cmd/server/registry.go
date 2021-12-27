package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	"github.com/calmato/shs-web/api/internal/lesson/api"
	db "github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type registry struct {
	lessonServer lesson.LessonServiceServer
}

type params struct {
	insecure            bool
	logger              *zap.Logger
	db                  *database.Client
	classroomServiceURL string
}

type gRPCClient struct {
	classroom classroom.ClassroomServiceClient
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
		Logger:           params.logger,
		WaitGroup:        &sync.WaitGroup{},
		Database:         db.NewDatabase(dbParams),
		ClassroomService: cli.classroom,
	}

	return &registry{
		lessonServer: api.NewLessonService(apiParams),
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

	return &gRPCClient{
		classroom: classroom.NewClassroomServiceClient(classroomConn),
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
