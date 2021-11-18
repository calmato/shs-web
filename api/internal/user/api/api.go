package api

import (
	"errors"
	"sync"
	"time"

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
}

type userService struct {
	user.UnimplementedUserServiceServer
	now         func() time.Time
	logger      *zap.Logger
	sharedGroup *singleflight.Group
	waitGroup   *sync.WaitGroup
	validator   validation.RequestValidation
}

func NewUserService(params *Params) user.UserServiceServer {
	return &userService{
		now:         jst.Now,
		logger:      params.Logger,
		waitGroup:   params.WaitGroup,
		sharedGroup: &singleflight.Group{},
		validator:   validation.NewRequestValidation(),
	}
}

func gRPCError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, validation.ErrRequestValidation):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
