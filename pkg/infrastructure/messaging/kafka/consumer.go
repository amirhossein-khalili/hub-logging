package kafka

import (
	"encoding/json"
	"hub_logging/pkg/application"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
	handlers map[string]func(application.DomainEvent)
}

func NewKafkaConsumer(brokers []string) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &KafkaConsumer{
		consumer: consumer,
		handlers: make(map[string]func(application.DomainEvent)),
	}, nil
}

func (c *KafkaConsumer) Subscribe(topic string, handler func(event application.DomainEvent)) error {
	c.handlers[topic] = handler
	return nil
}

func (c *KafkaConsumer) Start() error {
	for topic := range c.handlers {
		partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if err != nil {
			return err
		}
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				var event application.DomainEvent
				if err := json.Unmarshal(msg.Value, &event); err == nil {
					if handler, ok := c.handlers[topic]; ok {
						handler(event)
					}
				}
			}
		}(partitionConsumer)
	}
	return nil
}

func (c *KafkaConsumer) Close() error {
	return c.consumer.Close()
}
