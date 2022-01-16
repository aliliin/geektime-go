package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// NewConfig  Init config 初始化 viper
func NewConfig() (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)
	path, _ := os.Getwd()
	v.AddConfigPath(path + "/configs")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	return v, err
}
