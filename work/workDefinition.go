package work

import (
	"time"

	"github.com/google/uuid"
)

type workDefinition struct {
	Id          string
	WorkType    string
	Payload     string
	TimeCreated time.Time
}

func NewWorkDefinition(workType string, payload string) *workDefinition {
	return &workDefinition{
		Id:          uuid.NewString(),
		WorkType:    workType,
		Payload:     payload,
		TimeCreated: time.Now(),
	}
}
