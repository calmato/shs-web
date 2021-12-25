package cmd

import (
	"sync"

	"github.com/calmato/shs-web/api/internal/lesson/api"
	db "github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/proto/lesson"
	"go.uber.org/zap"
)

type registry struct {
	lessonServer lesson.LessonServiceServer
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
		lessonServer: api.NewLessonService(apiParams),
	}
}
