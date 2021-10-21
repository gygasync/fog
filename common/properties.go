package common

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Properties map[string]string

func LoadProperties(environment string) (Properties, error) {
	props := Properties{}

	propPath, err := filepath.Abs(filepath.Dir("../"))

	propPath = path.Join(propPath, "go.props")

	if err != nil {
		panic(err)
	}

	if environment != "" {
		propPath = path.Join(propPath, ".", environment)
	}

	file, err := os.Open(propPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := scanner.Text()
		if equal := strings.Index(entry, "="); equal >= 0 {
			if key := strings.TrimSpace(entry[:equal]); len(key) >= 0 {
				value := ""
				if len(entry) > equal {
					value = strings.TrimSpace(entry[equal+1:])
				}
				props[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return props, nil
}
