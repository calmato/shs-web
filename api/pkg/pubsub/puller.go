package pubsub

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/calmato/shs-web/api/pkg/backoff"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type puller struct {
	*pubsub.Subscription
	opts *PullerOptions
}

/**
 * options
 */
type PullerOptions struct {
	timeout                time.Duration
	concurrency            int
	logger                 *zap.Logger
	maxRetries             int64
	maxExtension           time.Duration
	maxExtensionPeriod     time.Duration
	maxOutstandingMessages int
	maxOutstandingBytes    int
	numGoroutines          int
	synchronous            bool
}

type PullerOption func(*PullerOptions)

func WithPullerTimeout(timeout time.Duration) PullerOption {
	return func(o *PullerOptions) {
		o.timeout = timeout
	}
}

func WithPullerConcurrency(concurrency int) PullerOption {
	return func(o *PullerOptions) {
		o.concurrency = concurrency
	}
}

func WithPullerLogger(logger *zap.Logger) PullerOption {
	return func(o *PullerOptions) {
		o.logger = logger
	}
}

func WithPullerMaxRetries(max int64) PullerOption {
	return func(o *PullerOptions) {
		o.maxRetries = max
	}
}

func WithPullerMaxExtension(max time.Duration) PullerOption {
	return func(o *PullerOptions) {
		o.maxExtension = max
	}
}

func WithPullerMaxExtensionPeriod(max time.Duration) PullerOption {
	return func(o *PullerOptions) {
		o.maxExtensionPeriod = max
	}
}

func WithPullerMaxOutstandingMessages(max int) PullerOption {
	return func(o *PullerOptions) {
		o.maxOutstandingMessages = max
	}
}

func WithPullerMaxOutstandingBytes(max int) PullerOption {
	return func(o *PullerOptions) {
		o.maxOutstandingBytes = max
	}
}

func WithPullerNumGoroutines(max int) PullerOption {
	return func(o *PullerOptions) {
		o.numGoroutines = max
	}
}

func WithPullerSync(sync bool) PullerOption {
	return func(o *PullerOptions) {
		o.synchronous = sync
	}
}

/**
 * main
 */
func NewPuller(cli *Client, subscriptionID string, opts ...PullerOption) Puller {
	subscription := cli.Client.Subscription(subscriptionID)
	dopts := &PullerOptions{}
	for i := range opts {
		opts[i](dopts)
	}
	subscription.ReceiveSettings = pubsub.ReceiveSettings{
		MaxExtension:           dopts.maxExtension,
		MaxExtensionPeriod:     dopts.maxExtensionPeriod,
		MaxOutstandingMessages: dopts.maxOutstandingMessages,
		MaxOutstandingBytes:    dopts.maxOutstandingBytes,
		NumGoroutines:          dopts.numGoroutines,
		Synchronous:            dopts.synchronous,
	}
	return &puller{
		Subscription: subscription,
		opts:         dopts,
	}
}

func (p *puller) Pull(ctx context.Context, msgCh chan<- *Message) error {
	retry := backoff.NewExponentialBackoff(p.opts.maxRetries)
	receive := func(ctx context.Context, msg *pubsub.Message) {
		m := &Message{
			ID:         msg.ID,
			Data:       msg.Data,
			Attributes: msg.Attributes,
			Ack:        msg.Ack,
			Nack:       msg.Nack,
		}
		select {
		case msgCh <- m:
		case <-ctx.Done():
			m.Nack()
		}
		retry.Reset()
	}

	var err error
	for retry.Continue() {
		select {
		case <-retry.Wait():
		case <-ctx.Done():
			return nil
		}
		err = p.Subscription.Receive(ctx, receive)
		if err == nil {
			return nil
		}
		switch {
		case status.Code(err) == codes.Canceled:
		default:
			p.opts.logger.Error("Failed to receive", zap.Error(err))
		}
		if isRetryable(err) {
			p.opts.logger.Info("Retry to receive from subscription")
			continue
		}
		return err
	}
	return err
}
