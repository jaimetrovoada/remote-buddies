package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	FRONTEND_URL string `mapstructure:"FRONTEND_URL"`
	DATABASE_URL string `mapstructure:"DATABASE_URL"`

	GITHUB_CLIENT_ID      string `mapstructure:"GITHUB_CLIENT_ID"`
	GITHUB_CLIENT_SECRET  string `mapstructure:"GITHUB_CLIENT_SECRET"`
	GITHUB_CLIENT_SCOPE   string `mapstructure:"GITHUB_CLIENT_SCOPE"`
	GITHUB_CLIENT_CALLBCK string `mapstructure:"GITHUB_CLIENT_CALLBCK"`

	SESSION_SECRET   string        `mapstructure:"SESSION_SECRET"`
	JWT_SECRET       string        `mapstructure:"JWT_SECRET"`
	TOKEN_EXPIRATION time.Duration `mapstructure:"TOKEN_EXPIRATION"`
	TOKEN_MAX_AGE    int           `mapstructure:"TOKEN_MAX_AGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
