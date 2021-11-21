package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port                 int64  `envconfig:"PORT" default:"8080"`
	MetricsPort          int64  `envconfig:"METRICS_PORT" default:"9090"`
	LogPath              string `envconfig:"LOG_PATH" default:""`
	LogLevel             string `envconfig:"LOG_LEVEL" default:"info"`
	GRPCInsecure         bool   `envconfig:"GRPC_INSECURE" default:"true"`
	GCPServiceKeyJSON    string `envconfig:"GCP_SERVICE_KEY_JSON" required:"true"`
	GCPStorageBucketName string `envconfig:"GCP_STORAGE_BUCKET_NAME" default:""`
	UserServiceURL       string `envconfig:"USER_SERVICE_URL" default:""`
}

func NewConfig() (*Config, error) {
	conf := &Config{}
	if err := envconfig.Process("", conf); err != nil {
		return conf, fmt.Errorf("config: failed to new config: %w", err)
	}
	return conf, nil
}
