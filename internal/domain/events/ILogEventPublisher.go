package events

import "hub_logging/internal/domain/aggregates"

// ILogEventPublisher is the Subject interface that publishes domain events.
type ILogEventPublisher interface {
	PublishLogCreated(logAgg aggregates.LogAggregate)
	Attach(observer ILogObserver)
	Detach(observer ILogObserver)
}
