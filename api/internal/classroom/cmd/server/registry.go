package cmd

import (
	"sync"

	"github.com/calmato/shs-web/api/internal/classroom/api"
	db "github.com/calmato/shs-web/api/internal/classroom/database"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/proto/classroom"
	"go.uber.org/zap"
)

type registry struct {
	classroomServer classroom.ClassroomServiceServer
}

type params struct {
	logger *zap.Logger
	db     *database.Client
}

func newRegistry(params *params) *registry {
	dbParams := &db.Params{
		Database: params.db,
	}
	apiParams := &api.Params{
		Logger:    params.logger,
		WaitGroup: &sync.WaitGroup{},
		Database:  db.NewDatabase(dbParams),
	}
	return &registry{
		classroomServer: api.NewClassroomService(apiParams),
	}
}
