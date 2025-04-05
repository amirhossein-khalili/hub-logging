package application

type DomainEvent interface {
	EventName() string
}

type EventPublisher interface {
	Publish(event DomainEvent) error
}

type EventConsumer interface {
	Subscribe(topic string, handler func(event DomainEvent)) error
	Start() error
	Close() error
}
