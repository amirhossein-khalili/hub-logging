// pkg/infrastructure/rabbitmq/consumer.go
package rabbitmq

import (
	"encoding/json"
	"hub_logging/pkg/application"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	handlers map[string]func(application.DomainEvent)
}

func NewRabbitMQConsumer(url string) (*RabbitMQConsumer, error) {
	conn, err := amqp.Dial(url) // e.g., "amqp://guest:guest@localhost:5672/"
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQConsumer{
		conn:     conn,
		channel:  channel,
		handlers: make(map[string]func(application.DomainEvent)),
	}, nil
}

func (c *RabbitMQConsumer) Subscribe(queue string, handler func(event application.DomainEvent)) error {
	// Declare the queue
	_, err := c.channel.QueueDeclare(
		queue, // Queue name
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

	c.handlers[queue] = handler
	return nil
}

func (c *RabbitMQConsumer) Start() error {
	for queue, handler := range c.handlers {
		msgs, err := c.channel.Consume(
			queue, // Queue name
			"",    // Consumer tag (auto-generated if empty)
			true,  // Auto-ack
			false, // Exclusive
			false, // No-local
			false, // No-wait
			nil,   // Arguments
		)
		if err != nil {
			return err
		}

		go func(q string, h func(application.DomainEvent)) {
			for msg := range msgs {
				var event application.DomainEvent
				if err := json.Unmarshal(msg.Body, &event); err == nil {
					h(event)
				}
			}
		}(queue, handler)
	}
	return nil
}

func (c *RabbitMQConsumer) Close() error {
	c.channel.Close()
	return c.conn.Close()
}
