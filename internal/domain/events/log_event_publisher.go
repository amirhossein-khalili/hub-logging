package events

import (
	"hub_logging/internal/domain/aggregates"
)

// LogEventPublisher is a concrete implementation of ILogEventPublisher.
type LogEventPublisher struct {
	observers []ILogObserver
}

// NewLogEventPublisher creates a new LogEventPublisher with an empty observers list.
func NewLogEventPublisher() *LogEventPublisher {
	return &LogEventPublisher{
		observers: make([]ILogObserver, 0),
	}
}

// PublishLogCreated notifies all observers about a newly created log.
func (p *LogEventPublisher) PublishLogCreated(logAgg aggregates.LogAggregate) {
	for _, obs := range p.observers {
		obs.OnLogCreated(logAgg)
	}
}

// Attach registers a new observer.
func (p *LogEventPublisher) Attach(observer ILogObserver) {
	p.observers = append(p.observers, observer)
}

// Detach unregisters an observer.
func (p *LogEventPublisher) Detach(observer ILogObserver) {
	for i, o := range p.observers {
		if o == observer {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}
