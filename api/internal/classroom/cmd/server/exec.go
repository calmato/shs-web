package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/calmato/shs-web/api/config/classroom/server"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/grpc"
	"github.com/calmato/shs-web/api/pkg/http"
	"github.com/calmato/shs-web/api/pkg/log"
	"github.com/calmato/shs-web/api/proto/classroom"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	ggrpc "google.golang.org/grpc"
)

//nolint:funlen
func Exec() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := config.NewConfig()
	if err != nil {
		return err
	}

	// Loggerの設定
	logParams := &log.Params{
		Path:  conf.LogPath,
		Level: conf.LogLevel,
	}
	logger, err := log.NewLogger(logParams)
	if err != nil {
		return err
	}

	// MySQLの設定
	dbParams := &database.Params{
		Socket:   conf.DBSocket,
		Host:     conf.DBHost,
		Port:     conf.DBPort,
		Database: conf.DBDatabase,
		Username: conf.DBUsername,
		Password: conf.DBPassword,
	}
	db, err := database.NewClient(dbParams)
	if err != nil {
		return err
	}

	// 依存関係の解決
	regParams := &params{
		insecure:       conf.GRPCInsecure,
		logger:         logger,
		db:             db,
		userServiceURL: conf.UserServiceURL,
	}
	reg, err := newRegistry(regParams)
	if err != nil {
		return err
	}

	// gRPC Serverの設定
	gRPCParams := &grpc.OptionParams{
		Logger: logger,
	}
	gRPCOpts := grpc.NewGRPCOptions(gRPCParams)

	s := ggrpc.NewServer(gRPCOpts...)
	classroom.RegisterClassroomServiceServer(s, reg.classroomServer)

	gs, err := grpc.NewGRPCServer(s, conf.Port)
	if err != nil {
		return err
	}

	// Metrics Serverの設定
	ms := http.NewMetricsServer(conf.MetricsPort)

	// Serverの起動
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		err = ms.Serve()
		if err != nil {
			logger.Error("Failed to run metrics server", zap.Error(err))
		}
		return
	})
	eg.Go(func() (err error) {
		err = gs.Serve()
		if err != nil {
			logger.Error("Failed to run gRPC server", zap.Error(err))
		}
		return
	})
	logger.Info("Started server", zap.Int64("port", conf.Port))

	// シグナル検知設定
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-ectx.Done():
		logger.Error("Done context", zap.Error(ectx.Err()))
	case <-signalCh:
		logger.Info("Received signal")
		delay := time.Duration(conf.ShutdownDelaySec) * time.Second
		logger.Info("Pre-shutdown", zap.String("delay", delay.String()))
		time.Sleep(delay)
	}

	logger.Info("Shutdown...")
	if err = ms.Stop(ectx); err != nil {
		return err
	}
	gs.Stop()
	return eg.Wait()
}
