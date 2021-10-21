package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadProperties(t *testing.T) {
	if _, err := os.Stat("../go.props"); os.IsNotExist(err) {
		t.Error("go.props does not exist")
	}

	props, err := LoadProperties("")
	if err != nil {
		t.Error("Could not load go.props")
	}

	for k, v := range props {
		fmt.Println(k, "=", v)
	}
}

func TestLoadPropertiesEnv(t *testing.T) {
	filename := "go.props.test"
	file, err := ioutil.TempFile("../", filename)
	if err != nil {
		t.Error("Could not create a temporary test prop file.")
	}
	tempfile := file.Name()

	file.Write([]byte("test=pass"))
	file.Close()

	props, err := LoadProperties(tempfile[len("fog/go.props"):])
	if err != nil {
		os.Remove("../" + filename)
		t.Error("Could load create a temporary test prop file.")
	}

	want := "pass"
	got := props["test"]

	if want != got {
		t.Errorf("got %s, wanted %s", got, want)
	}
	err = os.Remove(tempfile)
	if err != nil {
		t.Error("Could not remove file")
	}
}
