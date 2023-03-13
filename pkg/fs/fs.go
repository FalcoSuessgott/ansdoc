package fs

import (
	"io/ioutil"
	"log"
)

// ReadFile reads from a file.
func ReadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}

	return content
}
