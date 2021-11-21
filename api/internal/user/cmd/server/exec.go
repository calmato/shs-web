package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/calmato/shs-web/api/config/user/server"
	"github.com/calmato/shs-web/api/pkg/firebase"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/firebase/storage"
	"github.com/calmato/shs-web/api/pkg/grpc"
	"github.com/calmato/shs-web/api/pkg/http"
	"github.com/calmato/shs-web/api/pkg/log"
	"github.com/calmato/shs-web/api/proto/user"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
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

	// Firebaseの設定
	fbOpts := option.WithCredentialsJSON([]byte(conf.GCPServiceKeyJSON))
	fb, err := firebase.InitializeApp(ctx, nil, fbOpts)
	if err != nil {
		return err
	}

	// Firebase Authenticationの設定
	fa, err := authentication.NewClient(ctx, fb.App)
	if err != nil {
		return err
	}

	// Cloud Storageの設定
	gcs, err := storage.NewClient(ctx, fb.App, conf.GCPStorageBucketName)
	if err != nil {
		return err
	}

	// 依存関係の解決
	regParams := &params{
		logger:  logger,
		auth:    fa,
		storage: gcs,
	}
	reg := newRegistry(regParams)

	// gRPC Serverの設定
	gRPCParams := &grpc.OptionParams{
		Logger: logger,
	}
	gRPCOpts := grpc.NewGRPCServerOptions(gRPCParams)

	s := ggrpc.NewServer(gRPCOpts...)
	user.RegisterUserServiceServer(s, reg.userServer)

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
		delay := 20 * time.Second
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
