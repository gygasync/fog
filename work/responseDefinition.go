package work

import (
	"time"

	"github.com/google/uuid"
)

type responseDefinition struct {
	Id           string
	WorkId       string
	ResponseType string
	Payload      string
	TimeCreated  time.Time
}

func NewResponseDefinition(workDef workDefinition, responseType string, payload string) *responseDefinition {
	return &responseDefinition{
		Id:           uuid.NewString(),
		WorkId:       workDef.Id,
		ResponseType: responseType,
		Payload:      payload,
		TimeCreated:  time.Now(),
	}
}
