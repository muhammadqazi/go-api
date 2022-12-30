package core

import "github.com/spf13/viper"

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Port        string `mapstructure:"PORT"`
	DBUrl       string `mapstructure:"DB_URL"`
	SecretKey   string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./") // path to .env file
	viper.SetConfigName(c.Environment)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
