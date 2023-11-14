package kafka

import (
	"context"
	"github.com/raedmajeed/notification-service/pkg/config"
	"github.com/raedmajeed/notification-service/pkg/service"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafReaderStruct struct {
	reader *kafka.Reader
}

func NewKafkaReader() *KafReaderStruct {
	return &KafReaderStruct{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "test-topic",
			GroupID: "test-1",
		}),
	}
}

func (k *KafReaderStruct) EmailWriter(ctx context.Context, cfg *config.Conf) error {
	for {
		message, err := k.reader.FetchMessage(ctx)
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminating")
			return ctx.Err()
		default:
			log.Println("sending mail")
			//err := service.SendEmailToUser(email, k.conf.EMAIL, k.conf.PASSWORD)
			err := service.SendEmailToUser(message, string(cfg.EMAIL), string(cfg.PASSWORD))
			if err != nil {
				return err
			}
			err = k.CommitKafkaMessages(ctx, message)
			if err != nil {
				return err
			}
		}
	}
}

func (k *KafReaderStruct) CommitKafkaMessages(ctx context.Context, messages kafka.Message) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminated")
			return ctx.Err()
		default:
			log.Println("committing messages")
			err := k.reader.CommitMessages(ctx, messages)
			if err != nil {
				return err
			}
			return nil
		}
	}
}
