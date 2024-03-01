package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	config := new(Config)

	v := viper.New()
	v.AutomaticEnv()

	err := v.BindEnv("PORT")
	if err != nil {
		return nil, err
	}

	if err = v.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
