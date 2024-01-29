package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	EMAIL       string `mapStructure:"EMAIL"`
	PASSWORD    string `mapStructure:"PASSWORD"`
	KAFKABROKER string `mapstructure:"KAFKABROKER"`
	PORT        string `mapstructure:"PORT"`
}

func Configuration() (*Conf, error) {
	var conf Conf
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error", err)
		return &Conf{}, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return &Conf{}, err
	}

	return &conf, nil
}
