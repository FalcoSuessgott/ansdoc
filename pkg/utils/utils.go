package utils

import (
	"log"

	"gopkg.in/yaml.v3"
)

// FromYAML takes a yaml byte array and marshalls it into a map.
func FromYAML(b []byte, o interface{}) {
	if err := yaml.Unmarshal(b, &o); err != nil {
		log.Fatalf("cannot marshal %s to YAML: %v", string(b), err)
	}
}
