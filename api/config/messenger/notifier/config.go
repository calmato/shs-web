package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MetricsPort          int64  `envconfig:"METRICS_PORT" default:"9090"`
	LogPath              string `envconfig:"LOG_PATH" default:""`
	LogLevel             string `envconfig:"LOG_LEVEL" default:"info"`
	GRPCInsecure         bool   `envconfig:"GRPC_INSECURE" default:"true"`
	GCPProjectID         string `envconfig:"GCP_PROJECT_ID" default:""`
	GCPServiceKeyJSON    string `envconfig:"GCP_SERVICE_KEY_JSON" default:""`
	PubsubSubscriptionID string `envconfig:"PUBSUB_SUBSCRIPTION_ID" default:""`
	PubsubEmulatorHost   string `envconfig:"PUBSUB_EMULATOR_HOST" default:""`
	SendGridAPIKey       string `envconfig:"SENDGRID_API_KEY" default:""`
	SendGridTemplatePath string `envconfig:"SENDGRID_TEMPLATE_PATH" default:""`
	MailFromName         string `envconfig:"MAIL_FROM_NAME" default:""`
	MailFromAddress      string `envconfig:"MAIL_FROM_ADDRESS" default:""`
	UserServiceURL       string `envconfig:"User_SERVICE_URL" default:""`
	Concurrency          int64  `envconvig:"CONCURRENCY" default:"1"`
	MailboxCapacity      int64  `envconfig:"MAILBOX_CAPACITY" default:"1"`
	TeacherWebURL        string `envconfig:"TEACHER_WEB_URL" default:""`
	StudentWebURL        string `envconfig:"STUDENT_WEB_URL" default:""`
}

func NewConfig() (*Config, error) {
	conf := &Config{}
	if err := envconfig.Process("", conf); err != nil {
		return conf, fmt.Errorf("config: failed to new config: %w", err)
	}
	return conf, nil
}
