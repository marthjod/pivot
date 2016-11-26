package model

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	// "github.com/ghodss/yaml" not working correctly
)

type Environment string

type Runtime struct {
	CPU          string        `yaml:"cpu"`
	RAM          string        `yaml:"ram"`
	Disk         string        `yaml:"disk"`
	HostType     string        `yaml:"host_type"`
	NetworkZone  string        `yaml:"network_zone"`
	Environments []Environment `yaml:"environment"`
}

type Pivio struct {
	Id        string  `yaml:"id"`
	ShortName string  `yaml:"short_name"`
	Runtime   Runtime `yaml:"runtime"`
	// etc.
}

func FromFile(path string) (*Pivio, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pivio := Pivio{}
	if err := yaml.Unmarshal(data, &pivio); err != nil {
		return nil, err
	}

	return &pivio, nil
}
