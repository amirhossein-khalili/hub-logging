package aggregates

import (
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/valueobjects"
	"time"
)

// LogAggregate represents the aggregate root for log-related data and operations.
type LogAggregate struct {
	LogMessage entities.LogMessage
	Operations []entities.LogOperations
}

// NewLogAggregate creates a new LogAggregate with an initial LogMessage.
func NewLogAggregate(
	timestamp time.Time,
	statusCode valueobjects.StatusCode,
	httpMethod valueobjects.RequestMethod,
	routePath valueobjects.RoutePath,
	message string,
	userName string,
	destHostname string,
	sourceIP string,
) (*LogAggregate, error) {
	// Create the LogMessage entity using provided data
	logMessage := entities.LogMessage{
		Timestamp:    timestamp,
		StatusCode:   statusCode.Int(),
		HttpMethod:   httpMethod.String(),
		RoutePath:    routePath.String(),
		Message:      message,
		UserName:     userName,
		DestHostname: destHostname,
		SourceIP:     sourceIP,
	}

	// Return a new LogAggregate with an empty operations list
	return &LogAggregate{
		LogMessage: logMessage,
		Operations: []entities.LogOperations{},
	}, nil
}

// AddOperation adds a new operation to the log aggregate.
func (la *LogAggregate) AddOperation(operation string, performedBy string) {
	// Create a new LogOperations entity
	newOperation := entities.LogOperations{
		Operation:    operation,
		Timestamp:    time.Now(),
		LogMessageID: la.LogMessage.ID, // Assumes ID is set when persisted
		PerformedBy:  performedBy,
	}

	// Append the operation to the Operations slice
	la.Operations = append(la.Operations, newOperation)
}

// GetLogMessage returns the log message entity.
func (la *LogAggregate) GetLogMessage() entities.LogMessage {
	return la.LogMessage
}

// GetOperations returns the list of operations performed on the log.
func (la *LogAggregate) GetOperations() []entities.LogOperations {
	return la.Operations
}
