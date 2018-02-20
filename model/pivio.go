package model

import (
	"encoding/json"
	"io/ioutil"

	"io"

	"gopkg.in/yaml.v2"
)

type PivioEnvironment string

type Service struct {
	Name string `yaml:"service_name"`
}

type PivioRuntime struct {
	CPU          string             `yaml:"cpu"`
	RAM          string             `yaml:"ram"`
	Disk         string             `yaml:"disk"`
	HostType     string             `yaml:"host_type"`
	NetworkZone  string             `yaml:"network_zone"`
	Environments []PivioEnvironment `yaml:"environment"`
}

type PivioServices struct {
	Provides  []PivioServiceProvides `yaml:"provides"`
	DependsOn PivioServiceDependsOn  `yaml:"depends_on"`
}

type PivioServiceDependsOn struct {
	Internal []Service `yaml:"internal"`
	External []Service `yaml:"external"`
}

type PivioServiceProvides struct {
	Description       string `yaml:"description"`
	ServiceName       string `yaml:"service_name"`
	Protocol          string `yaml:"protocol"`
	Port              uint32 `yaml:"port"`
	TransportProtocol string `yaml:"transport_protocol"`
}

type Pivio struct {
	ID        string        `yaml:"id"`
	ShortName string        `yaml:"short_name"`
	Runtime   PivioRuntime  `yaml:"runtime"`
	Services  PivioServices `yaml:"service"`
	// etc.
}

func Read(r io.Reader) (*Pivio, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	pivio := Pivio{}
	if err := yaml.Unmarshal(data, &pivio); err != nil {
		return nil, err
	}

	return &pivio, nil
}

func ReadJSONMultiple(r io.Reader) ([]Pivio, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	pivios := []Pivio{}
	if err := json.Unmarshal(data, &pivios); err != nil {
		return nil, err
	}

	return pivios, nil
}
