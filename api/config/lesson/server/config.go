package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port                int64  `envconfig:"PORT" default:"8080"`
	MetricsPort         int64  `envconfig:"METRICS_PORT" default:"9090"`
	ShutdownDelaySec    int64  `envconfig:"SHUTDOWN_DELAY_SEC" default:"20"`
	LogPath             string `envconfig:"LOG_PATH" default:""`
	LogLevel            string `envconfig:"LOG_LEVEL" default:"info"`
	DBSocket            string `envconfig:"DB_SOCKET" default:"tcp"`
	DBHost              string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort              string `envconfig:"DB_PORT" default:"3306"`
	DBUsername          string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword          string `envconfig:"DB_PASSWORD" default:""`
	DBDatabase          string `envconfig:"DB_DATABASE" default:"lessons"`
	DBTimeZone          string `envconfig:"DB_TIMEZONE" default:""`
	GRPCInsecure        bool   `envconfig:"GRPC_INSECURE" default:"true"`
	ClassroomServiceURL string `envconfig:"CLASSROOM_SERVICE_URL" default:""`
	UserServiceURL      string `envconfig:"USER_SERVICE_URL" default:""`
	MessengerServiceURL string `envconfig:"MESSENGER_SERVICE_URL" default:""`
}

func NewConfig() (*Config, error) {
	conf := &Config{}
	if err := envconfig.Process("", conf); err != nil {
		return conf, fmt.Errorf("config: failed to new config: %w", err)
	}
	return conf, nil
}
