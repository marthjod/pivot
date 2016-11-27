package model

import (
	"gopkg.in/yaml.v2"
	// "github.com/ghodss/yaml" not working correctly
)

type OneTemplate struct {
	Memory int     `yaml:"memory"`
	VCPU   float32 `yaml:"vcpu"`
	Image  int     `yaml:"image"`
}

func (t *OneTemplate) ToYAML() ([]byte, error) {
	return yaml.Marshal(t)
}
