package events

import "hub_logging/internal/domain/aggregates"

// ILogObserver defines the Observer interface.
type ILogObserver interface {
	OnLogCreated(logAgg aggregates.LogAggregate)
}
