package work

import "encoding/json"

type errorDefinition struct {
	Err error
}

func NewErrorDefinition(err error) *errorDefinition {
	return &errorDefinition{Err: err}
}

func GenerateErrorPayload(err error) string {
	errBytes, err2 := json.Marshal(&errorDefinition{Err: err})
	if err2 != nil {
		panic(err2)
	}

	return string(errBytes)
}
