package common

import (
	"fmt"
	"testing"
)

func TestLoadProperties(t *testing.T) {
	props, err := LoadProperties("")
	if err != nil {
		t.Error("Could not load go.props")
	}

	for k, v := range props {
		fmt.Println(k, "=", v)
	}
}
