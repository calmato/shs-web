package cmd

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/calmato/shs-web/api/config/messenger/notifier"
	"github.com/calmato/shs-web/api/internal/messenger/mailer"
	"github.com/calmato/shs-web/api/pkg/http"
	"github.com/calmato/shs-web/api/pkg/log"
	"github.com/calmato/shs-web/api/pkg/pubsub"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v3"
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

	// SendGridの設定
	f, err := os.Open(conf.SendGridTemplatePath)
	if err != nil {
		return err
	}
	defer f.Close()

	var templateMap map[string]string
	d := yaml.NewDecoder(f)
	if err := d.Decode(&templateMap); err != nil {
		return err
	}

	mailParams := &mailer.Params{
		Logger:      logger,
		APIKey:      conf.SendGridAPIKey,
		FromName:    conf.MailFromName,
		FromAddress: conf.MailFromAddress,
		TemplateMap: templateMap,
	}
	mailer := mailer.NewClient(mailParams)

	// PubSubの設定
	psOpts := option.WithCredentialsJSON([]byte(conf.GCPServiceKeyJSON))
	ps, err := pubsub.NewClient(ctx, conf.GCPProjectID, psOpts)
	if err != nil {
		return err
	}
	puller := pubsub.NewPuller(
		ps, conf.PubsubSubscriptionID,
		pubsub.WithPullerTimeout(10*time.Second),
		pubsub.WithPullerMaxRetries(3),
		pubsub.WithPullerConcurrency(1),
		pubsub.WithPullerLogger(logger),
	)

	// Web URLの設定
	teacherWebURL, err := url.Parse(conf.TeacherWebURL)
	if err != nil {
		return err
	}
	studentWebURL, err := url.Parse(conf.StudentWebURL)
	if err != nil {
		return err
	}

	// 依存関係の解決
	regParams := &params{
		insecure:       conf.GRPCInsecure,
		logger:         logger,
		teacherWebURL:  teacherWebURL,
		studentWebURL:  studentWebURL,
		mailer:         mailer,
		userServiceURL: conf.UserServiceURL,
	}
	reg, err := newRegistry(regParams)
	if err != nil {
		return err
	}

	// Workerの設定
	n := reg.notifier

	// Metrics Serverの設定
	ms := http.NewMetricsServer(conf.MetricsPort)

	// Workerの起動
	eg, ectx := errgroup.WithContext(ctx)
	msgCh := make(chan *pubsub.Message, 1)
	eg.Go(func() (err error) {
		err = ms.Serve()
		if err != nil {
			logger.Error("Failed to run metrics server", zap.Error(err))
		}
		return
	})
	eg.Go(func() (err error) {
		err = puller.Pull(ectx, msgCh)
		if status.Code(err) == codes.Canceled {
			return nil
		}
		return
	})
	eg.Go(func() (err error) {
		err = n.Run(ectx, msgCh, int(conf.Concurrency), int(conf.MailboxCapacity))
		if err != nil {
			logger.Error("Failed to run notifier", zap.Error(err))
		}
		return
	})
	logger.Info("Started notifier")

	// シグナル検知設定
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-ectx.Done():
		logger.Error("Done context", zap.Error(ectx.Err()))
	case signal := <-signalCh:
		logger.Info("Received signal", zap.String("signal", signal.String()))
		cancel()
	}

	logger.Info("Shutdown...")
	if err = ms.Stop(ectx); err != nil {
		return err
	}
	return eg.Wait()
}
