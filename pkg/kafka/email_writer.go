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

func NewKafkaReader(cfg config.Conf) *KafReaderStruct {
	return &KafReaderStruct{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{cfg.KAFKABROKER},
			Topic:   "email-service-2",
			GroupID: "email-1",
		}),
	}
}

func (k *KafReaderStruct) EmailWriter(ctx context.Context, cfg *config.Conf) {
	msgCh := make(chan kafka.Message)
	go k.ReadMessage(ctx, cfg, msgCh)
	log.Println("message consumer listening")
	for {
		message, _ := k.reader.FetchMessage(ctx)
		if message.Value != nil {
			msgCh <- message
		}
	}
}

func (k *KafReaderStruct) ReadMessage(ctx context.Context, cfg *config.Conf, msgCh chan kafka.Message) {
	for {

		select {
		case <-ctx.Done():
			log.Fatalf("context termination")
		case message := <-msgCh:
			service.SendEmailToUser(message, string(cfg.EMAIL), string(cfg.PASSWORD))
			_ = k.CommitKafkaMessages(ctx, message)
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
