//go:generate mockgen -source=$GOFILE -package mock_$GOPACKAGE -destination=./../../mock/pkg/$GOPACKAGE/$GOFILE
package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	*pubsub.Client
}

type Message struct {
	ID         string
	Data       []byte
	Attributes map[string]string
	Ack        func()
	Nack       func()
}

func NewClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*Client, error) {
	cli, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{Client: cli}, nil
}

type Publisher interface {
	Publish(ctx context.Context, msg *Message) (string, error)
	Stop()
}

type Puller interface {
	Pull(ctx context.Context, msgCh chan<- *Message) error
}

/**
 * common private methods
 */
func isRetryable(err error) bool {
	if err == nil {
		return false
	}
	if err == context.DeadlineExceeded {
		return true
	}
	s, ok := status.FromError(err)
	if !ok {
		return false
	}
	switch s.Code() {
	case codes.Canceled, codes.DeadlineExceeded:
		return true
	case codes.Internal, codes.Unavailable:
		return true
	default:
		return false
	}
}
