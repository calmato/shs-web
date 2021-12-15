package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/calmato/shs-web/api/config/gateway/teacher"
	"github.com/calmato/shs-web/api/pkg/cors"
	"github.com/calmato/shs-web/api/pkg/firebase"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/http"
	"github.com/calmato/shs-web/api/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
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

	// 依存関係の解決
	regParams := &params{
		insecure:       conf.GRPCInsecure,
		logger:         logger,
		auth:           fa,
		userServiceURL: conf.UserServiceURL,
	}
	reg, err := newRegistry(regParams)
	if err != nil {
		return err
	}

	// HTTP Serverの設定
	httpOpts := []gin.HandlerFunc{gin.Recovery()}

	cm := cors.NewGinMiddleware()
	httpOpts = append(httpOpts, cm)

	lm, err := log.NewGinMiddleware(logParams)
	if err != nil {
		return err
	}
	httpOpts = append(httpOpts, lm)

	rt := newRouter(reg, httpOpts...)
	hs := http.NewHTTPServer(rt, conf.Port)

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
		err = hs.Serve()
		if err != nil {
			logger.Error("Failed to run http server", zap.Error(err))
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
	if err = hs.Stop(ectx); err != nil {
		return err
	}
	if err = ms.Stop(ectx); err != nil {
		return err
	}
	return eg.Wait()
}
