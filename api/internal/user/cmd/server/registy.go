package cmd

import (
	"sync"

	"github.com/calmato/shs-web/api/internal/user/api"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/firebase/storage"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
)

type registry struct {
	userServer user.UserServiceServer
}

type params struct {
	logger  *zap.Logger
	auth    authentication.Client
	storage storage.Client
}

func newRegistry(params *params) *registry {
	apiParams := &api.Params{
		Logger:    params.logger,
		WaitGroup: &sync.WaitGroup{},
	}
	return &registry{
		userServer: api.NewUserService(apiParams),
	}
}
