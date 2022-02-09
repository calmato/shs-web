package api

import (
	"errors"
	"sync"
	"time"

	"github.com/calmato/shs-web/api/internal/messenger/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/pubsub"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/messenger"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Params struct {
	Logger        *zap.Logger
	WaitGroup     *sync.WaitGroup
	Publisher     pubsub.Publisher
	UserService   user.UserServiceClient
	LessonService lesson.LessonServiceClient
}

type messengerService struct {
	messenger.UnimplementedMessengerServiceServer
	now         func() time.Time
	logger      *zap.Logger
	sharedGroup *singleflight.Group
	waitGroup   *sync.WaitGroup
	validator   validation.RequestValidation
	publisher   pubsub.Publisher
	user        user.UserServiceClient
	lesson      lesson.LessonServiceClient
}

func NewMessengerService(params *Params) messenger.MessengerServiceServer {
	return &messengerService{
		now:         jst.Now,
		logger:      params.Logger,
		waitGroup:   params.WaitGroup,
		sharedGroup: &singleflight.Group{},
		validator:   validation.NewRequestValidation(),
		publisher:   params.Publisher,
		user:        params.UserService,
		lesson:      params.LessonService,
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
	case errors.Is(err, validation.ErrRequestValidation):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
