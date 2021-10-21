package files

import "testing"

func TestAddValidPathToRegisteredPaths(t *testing.T) {
	RegisterPath("../files")
	RegisterPath("../files2")
	RegisterPath("../learning")
	got := DumpPaths()
	want := "../files\n../learning"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
