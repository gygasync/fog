package work

import "github.com/google/uuid"

type workDefinition struct {
	id       string
	workType string
	payload  string
}

func NewWorkDefinition(workType string, payload string) *workDefinition {
	return &workDefinition{
		id:       uuid.NewString(),
		workType: workType,
		payload:  payload,
	}
}
