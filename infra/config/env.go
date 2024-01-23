package config

import "github.com/spf13/viper"

var env *envconfig

type envconfig struct {
	DBUrl           string `mapstructure:"DATABASE_URL"`
	PORT            string `mapstructure:"PORT"`
	ENV             string `mapstructure:"ENVIROMENT"`
	MIGRATIONS_PATH string `mapstructure:"MIGRATIONS_PATH"`
}

func GetEnvConfig() *envconfig {
	return env
}

func LoadEnv(path string) (*envconfig, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	return env, nil
}
