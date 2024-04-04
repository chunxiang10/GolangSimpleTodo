package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Server struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}

type config struct {
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
}

var conf = new(config)

func Conf() *config {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s", err))
	}
	return conf
}
