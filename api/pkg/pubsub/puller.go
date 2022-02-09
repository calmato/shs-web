package pubsub

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/calmato/shs-web/api/pkg/backoff"
	"go.uber.org/zap"
)

type puller struct {
	*pubsub.Subscription
	opts *PullerOptions
}

/**
 * options
 */
type PullerOptions struct {
	timeout     time.Duration
	maxRetries  int64
	concurrency int
	logger      *zap.Logger
}

type PullerOption func(*PullerOptions)

func WithPullerTimeout(timeout time.Duration) PullerOption {
	return func(o *PullerOptions) {
		o.timeout = timeout
	}
}

func WithPullerMaxRetries(max int64) PullerOption {
	return func(o *PullerOptions) {
		o.maxRetries = max
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

/**
 * main
 */
func NewPuller(cli *Client, subscriptionID string, opts ...PullerOption) Puller {
	subscription := cli.Client.Subscription(subscriptionID)
	dopts := &PullerOptions{}
	for i := range opts {
		opts[i](dopts)
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
		if p.opts.logger != nil {
			p.opts.logger.Error("Failed to receive", zap.Error(err))
		}
		if !isRetryable(err) {
			return err
		}
		if p.opts.logger != nil {
			p.opts.logger.Info("Retry to receive from subscription")
		}
	}
	return err
}
