package main

import (
	"context"
	"github.com/raedmajeed/notification-service/pkg/config"
	"github.com/raedmajeed/notification-service/pkg/kafka"
	"log"
	"os"
	"os/signal"
)

func main() {
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt)
	cfg, err := config.Configuration()
	log.Println("VERSION:3 -> ", cfg)
	if err != nil {
		log.Fatalf("unable to load config file, aborting")
	}
	reader := kafka.NewKafkaReader(*cfg)
	go reader.EmailWriter(context.Background(), cfg)
	if err != nil {
		return
	}
	<-sign
	log.Println("reading from kafka complete")
}
