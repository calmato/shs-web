package cmd

import (
	"sync"

	"github.com/calmato/shs-web/api/internal/user/api"
	db "github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/pkg/database"
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
	db      *database.Client
	storage storage.Client
}

func newRegistry(params *params) *registry {
	dbParams := &db.Params{
		Database: params.db,
		Auth:     params.auth,
	}
	apiParams := &api.Params{
		Logger:    params.logger,
		WaitGroup: &sync.WaitGroup{},
		Database:  db.NewDatabase(dbParams),
	}
	return &registry{
		userServer: api.NewUserService(apiParams),
	}
}
