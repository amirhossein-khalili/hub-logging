package kafka

import (
	"encoding/json"
	"hub_logging/pkg/application"

	"github.com/IBM/sarama"
)

type KafkaPublisher struct {
	producer sarama.SyncProducer
}

func NewKafkaPublisher(brokers []string) (*KafkaPublisher, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all replicas to acknowledge
	config.Producer.Retry.Max = 5                    // Retry up to 5 times
	config.Producer.Return.Successes = true          // Ensure we get confirmation of success

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &KafkaPublisher{producer: producer}, nil
}

func (p *KafkaPublisher) Publish(event application.DomainEvent) error {
	// Serialize event to JSON
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Publish to Kafka
	topic := "events" // You can make this configurable
	_, _, err = p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	})
	return err
}

func (p *KafkaPublisher) Close() error {
	return p.producer.Close()
}
