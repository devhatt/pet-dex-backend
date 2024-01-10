package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	DBUrl string `mapstructure:"DATABASE_URL"`
	PORT  string `mapstructure:"PORT"`
	ENV   string `mapstructure:"ENVIROMENT"`
}

func GetConfig() *config {
	return cfg
}

func LoadConfig(path string) (*config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
