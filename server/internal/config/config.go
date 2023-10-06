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
	PORT             string        `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {

	viper.BindEnv("PORT")

	viper.BindEnv("FRONTEND_URL")
	viper.BindEnv("DATABASE_URL")

	viper.BindEnv("SESSION_SECRET")
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("TOKEN_EXPIRATION")
	viper.BindEnv("TOKEN_MAX_AGE")

	viper.BindEnv("GITHUB_CLIENT_ID")
	viper.BindEnv("GITHUB_CLIENT_SECRET")
	viper.BindEnv("GITHUB_CLIENT_SCOPE")
	viper.BindEnv("GITHUB_CLIENT_CALLBCK")
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}
