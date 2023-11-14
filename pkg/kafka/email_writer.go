package kafka

import (
	"context"
	"github.com/raedmajeed/notification-service/pkg/service"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafReaderStruct struct {
	reader *kafka.Reader
}

func NewKafkaReader() *KafReaderStruct {
	return &KafReaderStruct{
		kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{"localhost:7070"},
			GroupID:  "group-test",
			MaxBytes: 10e1,
		}),
	}
}

func (k *KafReaderStruct) EmailWriter(ctx context.Context) error {
	for {
		messages := make(chan kafka.Message)
		message, err := k.reader.FetchMessage(ctx)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminating")
			return ctx.Err()
		case messages <- message:
			err := service.SendEmailToUser(ctx, messages)
			if err != nil {
				return err
			}
			err = k.CommitKafkaMessages(ctx, messages)
			if err != nil {
				return err
			}
		}
	}
}

func (k *KafReaderStruct) CommitKafkaMessages(ctx context.Context, messages chan kafka.Message) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminated")
		case message := <-messages:
			err := k.reader.CommitMessages(ctx, message)
			if err != nil {
				return err
			}
		}
	}
}
