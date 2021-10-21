package files

import (
	"log"
	"os"
	"strings"
)

var registeredPaths []string

func RegisterPath(path string) {
	dir, err := os.Stat(path)
	if err != nil || !dir.IsDir() {
		log.Println("Path " + path + " is not valid")
		return
	}
	registeredPaths = append(registeredPaths, path)
}

func DumpPaths() string {
	return strings.Join(registeredPaths, "\n")
}
