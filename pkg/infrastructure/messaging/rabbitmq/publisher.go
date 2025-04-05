// pkg/infrastructure/rabbitmq/publisher.go
package rabbitmq

import (
	"encoding/json"
	"hub_logging/pkg/application"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewRabbitMQPublisher(url, queue string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(url) // e.g., "amqp://guest:guest@localhost:5672/"
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	// Declare the queue
	_, err = channel.QueueDeclare(
		queue, // Queue name
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQPublisher{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}

func (p *RabbitMQPublisher) Publish(event application.DomainEvent) error {
	// Serialize event to JSON
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Publish to RabbitMQ
	err = p.channel.Publish(
		"",      // Exchange (empty for default)
		p.queue, // Routing key (queue name)
		false,   // Mandatory
		false,   // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	return err
}

func (p *RabbitMQPublisher) Close() error {
	p.channel.Close()
	return p.conn.Close()
}
