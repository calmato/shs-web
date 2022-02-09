package pubsub

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/calmato/shs-web/api/pkg/backoff"
	"go.uber.org/zap"
)

type publisher struct {
	*pubsub.Topic
	opts *PublisherOptions
}

/**
 * options
 */
type PublisherOptions struct {
	timeout    time.Duration
	maxRetries int64
	logger     *zap.Logger
}

type PublisherOption func(*PublisherOptions)

func WithPublisherTimeout(timeout time.Duration) PublisherOption {
	return func(o *PublisherOptions) {
		o.timeout = timeout
	}
}

func WithPublisherMaxRetries(max int64) PublisherOption {
	return func(o *PublisherOptions) {
		o.maxRetries = max
	}
}

func WithPublisherLogger(logger *zap.Logger) PublisherOption {
	return func(o *PublisherOptions) {
		o.logger = logger
	}
}

/**
 * main
 */
func NewPublisher(cli *Client, topicID string, opts ...PublisherOption) Publisher {
	topic := cli.Client.Topic(topicID)
	dopts := &PublisherOptions{}
	for i := range opts {
		opts[i](dopts)
	}
	return &publisher{
		Topic: topic,
		opts:  dopts,
	}
}

func (p *publisher) Publish(ctx context.Context, msg *Message) (string, error) {
	if p.opts.maxRetries == 0 {
		return p.syncPublish(ctx, msg)
	}
	duration := time.Millisecond * 100
	retry := backoff.NewFixedIntervalBackoff(duration, p.opts.maxRetries)
	var serverID string
	var err error
	for retry.Continue() {
		serverID, err = p.syncPublish(ctx, msg)
		if err == nil {
			return serverID, nil
		}
		if !isRetryable(err) {
			return serverID, err
		}
		if p.opts.logger != nil {
			p.opts.logger.Warn("Retry to publish from topic", zap.Error(err))
		}

		select {
		case <-retry.Wait():
			continue
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	return serverID, err
}

func (p *publisher) syncPublish(ctx context.Context, msg *Message) (string, error) {
	if timeout := p.opts.timeout; timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	return p.publish(ctx, msg).Get(ctx)
}

func (p *publisher) publish(ctx context.Context, msg *Message) *pubsub.PublishResult {
	params := &pubsub.Message{
		Data:       msg.Data,
		Attributes: msg.Attributes,
	}
	return p.Topic.Publish(ctx, params)
}

func (p *publisher) Stop() {
	p.Topic.Stop()
}
