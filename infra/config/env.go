package config

import "github.com/spf13/viper"

var env *Envconfig

type Envconfig struct {
	DB_USER                    string `mapstructure:"DB_USER"`
	DB_PASSWORD                string `mapstructure:"DB_PASSWORD"`
	DB_DATABASE                string `mapstructure:"DB_DATABASE"`
	DB_HOST                    string `mapstructure:"DB_HOST"`
	DB_PORT                    string `mapstructure:"DB_PORT"`
	API_PORT                   string `mapstructure:"API_PORT"`
	ENV                        string `mapstructure:"ENVIRONMENT"`
	MIGRATIONS_PATH            string `mapstructure:"MIGRATIONS_PATH"`
	JWT_SECRET                 string `mapstructure:"JWT_SECRET"`
	GOOGLE_OAUTH_CLIENT_ID     string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GOOGLE_OAUTH_CLIENT_SECRET string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GOOGLE_REDIRECT_URL        string `mapstructure:"GOOGLE_REDIRECT_URL"`
	FACEBOOK_APP_ID            string `mapstructure:"FACEBOOK_APP_ID"`
	FACEBOOK_APP_SECRET        string `mapstructure:"FACEBOOK_APP_SECRET"`
}

func GetEnvConfig() *Envconfig {
	return env
}

func LoadEnv(path string) (*Envconfig, error) {
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
