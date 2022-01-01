package api

import (
	"errors"
	"sync"
	"time"

	"github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/internal/user/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Params struct {
	Logger    *zap.Logger
	WaitGroup *sync.WaitGroup
	Database  *database.Database
}

type userService struct {
	user.UnimplementedUserServiceServer
	now         func() time.Time
	logger      *zap.Logger
	sharedGroup *singleflight.Group
	waitGroup   *sync.WaitGroup
	db          *database.Database
	validator   validation.RequestValidation
}

func NewUserService(params *Params) user.UserServiceServer {
	return &userService{
		now:         jst.Now,
		logger:      params.Logger,
		waitGroup:   params.WaitGroup,
		sharedGroup: &singleflight.Group{},
		db:          params.Database,
		validator:   validation.NewRequestValidation(),
	}
}

func gRPCError(err error) error {
	if err == nil {
		return nil
	}

	if _, ok := status.FromError(err); ok {
		return err
	}

	switch {
	case errors.Is(err, validation.ErrRequestValidation),
		errors.Is(err, database.ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, database.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, database.ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, database.ErrNotImplemented):
		return status.Error(codes.Unimplemented, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
