package main

import (
	"context"
	"github.com/raedmajeed/notification-service/pkg/kafka"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	reader := kafka.NewKafkaReader()
	group, ctx := errgroup.WithContext(context.Background())

	//readerChan := make(chan kafka.KafReaderStruct)
	group.Go(func() error {
		return reader.EmailWriter(ctx)
	})

	err := group.Wait()
	if err != nil {
		return
	}

	log.Println("reading from kafka complete")
}
