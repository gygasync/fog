package definition

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	Id           string
	WorkId       string
	ResponseType string
	Payload      string
	TimeCreated  time.Time
}

func NewResponseDefinition(workDef Work, responseType string, payload string) *Response {
	return &Response{
		Id:           uuid.NewString(),
		WorkId:       workDef.Id,
		ResponseType: responseType,
		Payload:      payload,
		TimeCreated:  time.Now(),
	}
}
