package core

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment            string `mapstructure:"ENVIRONMENT"`
	Port                   string `mapstructure:"PORT"`
	DBUrl                  string `mapstructure:"DB_URL"`
	SecretKey              string `mapstructure:"JWT_SECRET"`
	ResetPasswordSecretKey string `mapstructure:"RESET_PASSWORD_JWT_SECRET"`
	JWTAlgorithm           string `mapstructure:"JWT_ALGORITHM"`
	ResetPasswordAlgorithm string `mapstructure:"RESET_PASSWORD_JWT_ALGORITHM"`
}

func LoadConfig() (c Config, err error) {

	viper.AddConfigPath("./") // path to .env file
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
