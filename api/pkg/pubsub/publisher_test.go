package pubsub

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newTestPublisher(ctx context.Context) (*publisher, error) {
	client, err := NewClient(ctx, "project-test")
	if err != nil {
		return nil, err
	}
	topic, err := client.Client.CreateTopic(ctx, testTopic)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		return nil, err
	}
	if status.Code(err) == codes.AlreadyExists {
		topic = client.Client.Topic(testTopic)
	}
	opts := &PublisherOptions{
		timeout:    time.Second,
		maxRetries: 2,
	}
	return &publisher{
		Topic: topic,
		opts:  opts,
	}, nil
}

func TestPublisher(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := NewClient(ctx, "project-test")
	require.NoError(t, err)
	defer client.Close()

	opts := []PublisherOption{
		WithPublisherTimeout(time.Second * 2),
		WithPublisherMaxRetries(2),
		WithPublisherLogger(zap.NewNop()),
	}
	publisher := NewPublisher(client, testTopic, opts...)
	assert.NotNil(t, publisher)
}

func TestPublisher_Publish(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t.Run("success publish", func(t *testing.T) {
		publisher, err := newTestPublisher(ctx)
		require.NoError(t, err)
		publisher.opts.maxRetries = 2
		defer publisher.Stop()

		params := &Message{
			Data:       []byte("hello"),
			Attributes: map[string]string{"key": "value"},
		}
		serverID, err := publisher.Publish(ctx, params)
		assert.NoError(t, err)
		assert.NotEmpty(t, serverID)
	})

	t.Run("success async publish", func(t *testing.T) {
		publisher, err := newTestPublisher(ctx)
		require.NoError(t, err)
		publisher.opts.maxRetries = 0
		defer publisher.Stop()

		params := &Message{
			Data:       []byte("hello"),
			Attributes: map[string]string{"key": "value"},
		}
		serverID, err := publisher.Publish(ctx, params)
		assert.NoError(t, err)
		assert.NotEmpty(t, serverID)
	})
}
