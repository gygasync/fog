package work

import "github.com/google/uuid"

type workDefinition struct {
	Id       string
	WorkType string
	Payload  string
}

func NewWorkDefinition(workType string, payload string) *workDefinition {
	return &workDefinition{
		Id:       uuid.NewString(),
		WorkType: workType,
		Payload:  payload,
	}
}
