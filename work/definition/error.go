package definition

import "encoding/json"

type Error struct {
	Err error
}

func NewErrorDefinition(err error) *Error {
	return &Error{Err: err}
}

func GenerateErrorPayload(err error) string {
	errBytes, err2 := json.Marshal(&Error{Err: err})
	if err2 != nil {
		panic(err2)
	}

	return string(errBytes)
}
