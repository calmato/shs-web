package pubsub

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newTestPuller(ctx context.Context) (*puller, error) {
	client, err := NewClient(ctx, "project-test")
	if err != nil {
		return nil, err
	}
	topic := client.Client.Topic(testTopic)
	conf := pubsub.SubscriptionConfig{Topic: topic}
	subscription, err := client.Client.CreateSubscription(ctx, testTopic, conf)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		return nil, err
	}
	if status.Code(err) == codes.AlreadyExists {
		subscription = client.Client.Subscription(testTopic)
	}
	opts := &PullerOptions{
		timeout:     time.Second,
		maxRetries:  2,
		concurrency: 1,
	}
	return &puller{
		Subscription: subscription,
		opts:         opts,
	}, nil
}

func TestPuller(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := NewClient(ctx, "project-test")
	require.NoError(t, err)
	defer client.Close()

	opts := []PullerOption{
		WithPullerTimeout(time.Second * 2),
		WithPullerMaxRetries(2),
		WithPullerConcurrency(2),
		WithPullerLogger(zap.NewNop()),
	}
	puller := NewPuller(client, testTopic, opts...)
	assert.NotNil(t, puller)
}

func TestPuller_Pull(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	publisher, err := newTestPublisher(ctx)
	require.NoError(t, err)
	defer publisher.Stop()

	msg := &Message{
		Data:       []byte("hello"),
		Attributes: map[string]string{"key": "value"},
	}

	t.Run("pull", func(t *testing.T) {
		puller, err := newTestPuller(ctx)
		require.NoError(t, err)

		msgCh := make(chan *Message, 1)
		eg := errgroup.Group{}
		eg.Go(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return puller.Pull(ctx, msgCh)
		})
		_, err = publisher.Publish(ctx, msg)
		assert.NoError(t, err)
		for i := 0; i < 1; i++ {
			m := <-msgCh
			m.Ack()
			assert.Equal(t, msg.Data, m.Data)
		}
		assert.NoError(t, eg.Wait())
	})
}
