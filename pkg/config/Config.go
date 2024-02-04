package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Conf struct {
	EMAIL       string `mapStructure:"EMAIL"`
	PASSWORD    string `mapStructure:"PASSWORD"`
	KAFKABROKER string `mapstructure:"KAFKABROKER"`
	PORT        string `mapstructure:"PORT"`
}

func Configuration() (*Conf, error) {
	var cfg Conf

	if err := godotenv.Load(".env"); err != nil {
		os.Exit(1)
	}
	cfg.PORT = os.Getenv("PORT")
	cfg.KAFKABROKER = os.Getenv("KAFKABROKER")
	cfg.EMAIL = os.Getenv("EMAIL")
	cfg.PASSWORD = os.Getenv("PASSWORD")

	log.Println(cfg)

	return &cfg, nil
}
