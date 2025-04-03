package mappers

import (
	model "hub_logging/external/infra/models"
	domain "hub_logging/internal/domain/entities"
)

// ToModelLogMessage converts a domain LogMessage into its DB model representation.
func ToModelLogMessage(entity domain.LogMessage) model.LogMessage {
	return model.LogMessage{
		ID:           entity.ID,
		Timestamp:    entity.Timestamp,
		StatusCode:   entity.StatusCode,
		HttpMethod:   entity.HttpMethod,
		RoutePath:    entity.RoutePath,
		Message:      entity.Message,
		UserName:     entity.UserName,
		DestHostname: entity.DestHostname,
		SourceIP:     entity.SourceIP,
	}
}

// ToEntityLogMessage converts a DB model LogMessage into the domain entity.
func ToEntityLogMessage(model model.LogMessage) domain.LogMessage {
	return domain.LogMessage{
		ID:           model.ID,
		Timestamp:    model.Timestamp,
		StatusCode:   model.StatusCode,
		HttpMethod:   model.HttpMethod,
		RoutePath:    model.RoutePath,
		Message:      model.Message,
		UserName:     model.UserName,
		DestHostname: model.DestHostname,
		SourceIP:     model.SourceIP,
	}
}
