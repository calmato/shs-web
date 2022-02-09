//go:generate mockgen -source=$GOFILE -package mock_$GOPACKAGE -destination=./../../../mock/messenger/$GOPACKAGE/$GOFILE
package mailer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/sendgrid/sendgrid-go"
	"go.uber.org/zap"
)

var (
	ErrInvalidArgument  = errors.New("mail: invalid argument")
	ErrUnauthenticated  = errors.New("mail: unauthenticated")
	ErrPermissionDenied = errors.New("mail: permission denied")
	ErrPayloadTooLong   = errors.New("mail: payload too long")
	ErrNotFound         = errors.New("mail: not found")
	ErrInternal         = errors.New("mail: internal")
	ErrUnavailable      = errors.New("mail: unavailable")
	ErrTimeout          = errors.New("mail: timeout")
	ErrUnknown          = errors.New("mail: unknown")
)

type Client interface {
	SendFromInfo(ctx context.Context, emailID, toName, toAddress string, substitutions map[string]interface{}) error
	MultiSend(ctx context.Context, emailID, fromName, fromAddress string, ps []*Personalization) error
	MultiSendFromInfo(ctx context.Context, emailID string, ps []*Personalization) error
}

type Params struct {
	Logger      *zap.Logger
	APIKey      string
	FromName    string
	FromAddress string
	TemplateMap map[string]string
}

type client struct {
	now         func() time.Time
	client      *sendgrid.Client
	logger      *zap.Logger
	fromName    string
	fromAddress string
	templateMap map[string]string
}

// NewClient - メール送信用クライアントの生成
func NewClient(params *Params) Client {
	return &client{
		now:         jst.Now,
		client:      sendgrid.NewSendClient(params.APIKey),
		logger:      params.Logger,
		fromName:    params.FromName,
		fromAddress: params.FromAddress,
		templateMap: params.TemplateMap,
	}
}

/**
 * private method
 */
func mailError(e error) error {
	if e == nil {
		return nil
	}

	switch {
	case errors.Is(e, context.Canceled):
		return fmt.Errorf("%w: %s", ErrTimeout, e.Error())
	case errors.Is(e, context.DeadlineExceeded):
		return fmt.Errorf("%w: %s", ErrTimeout, e.Error())
	}

	err, ok := e.(*SendGridError)
	if !ok {
		return fmt.Errorf("%w: %s", ErrUnknown, e.Error())
	}

	switch err.Code {
	case http.StatusBadRequest:
		return fmt.Errorf("%w: %s", ErrInvalidArgument, err.Error())
	case http.StatusUnauthorized:
		return fmt.Errorf("%w: %s", ErrUnauthenticated, err.Error())
	case http.StatusForbidden:
		return fmt.Errorf("%w: %s", ErrPermissionDenied, err.Error())
	case http.StatusRequestEntityTooLarge:
		return fmt.Errorf("%w: %s", ErrPayloadTooLong, err.Error())
	case http.StatusNotFound:
		return fmt.Errorf("%w: %s", ErrNotFound, err.Error())
	case http.StatusInternalServerError:
		return fmt.Errorf("%w: %s", ErrInternal, err.Error())
	case http.StatusBadGateway:
		return fmt.Errorf("%w: %s", ErrUnavailable, err.Error())
	case http.StatusGatewayTimeout:
		return fmt.Errorf("%w: %s", ErrTimeout, err.Error())
	default:
		return fmt.Errorf("%w: %s", ErrUnknown, err.Error())
	}
}
