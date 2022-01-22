package definition

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Work struct {
	Id          string
	WorkType    string
	Payload     string
	TimeCreated time.Time
}

func NewWorkDefinition(workType string, payload interface{}) *Work {
	msg, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return &Work{
		Id:          uuid.NewString(),
		WorkType:    workType,
		Payload:     string(msg),
		TimeCreated: time.Now(),
	}
}
