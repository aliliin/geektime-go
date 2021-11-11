package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpAddress string
	GrpcAddress string
}

func InitConfig() (*viper.Viper, error) {
	viper.AddConfigPath("./../week4/config/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {

		return nil, err
	}
	return viper.GetViper(), nil
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		HttpAddress: ":" + v.Sub("http").GetString("port"),
		GrpcAddress: ":" + v.Sub("grpc").GetString("port"),
	}
}
