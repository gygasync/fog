package work

import "github.com/google/uuid"

type responseDefinition struct {
	Id           string
	WorkId       string
	ResponseType string
	Payload      string
}

func NewResponseDefinition(workDef workDefinition, responseType string, payload string) *responseDefinition {
	return &responseDefinition{
		Id:           uuid.NewString(),
		WorkId:       workDef.Id,
		ResponseType: responseType,
		Payload:      payload,
	}
}
