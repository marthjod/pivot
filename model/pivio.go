package model

import (
	"encoding/json"
	"io/ioutil"

	"io"

	"gopkg.in/yaml.v2"
)

type PivioEnvironment string

type Service struct {
	Name string `yaml:"service_name" json:"service_name"`
}

type PivioRuntime struct {
	CPU          string             `yaml:"cpu" json:"cpu"`
	RAM          string             `yaml:"ram" json:"ram"`
	Disk         string             `yaml:"disk" json:"disk"`
	HostType     string             `yaml:"host_type" json:"host_type"`
	NetworkZone  string             `yaml:"network_zone" json:"network_zone"`
	Environments []PivioEnvironment `yaml:"environment" json:"environment"`
}

type PivioServices struct {
	Provides  []PivioServiceProvides `yaml:"provides" json:"provides"`
	DependsOn PivioServiceDependsOn  `yaml:"depends_on" json:"depends_on"`
}

type PivioServiceDependsOn struct {
	Internal []Service `yaml:"internal" json:"internal"`
	External []Service `yaml:"external" json:"external"`
}

type PivioServiceProvides struct {
	Description       string `yaml:"description" json:"description"`
	ServiceName       string `yaml:"service_name" json:"service_name"`
	Protocol          string `yaml:"protocol" json:"protocol"`
	Port              uint32 `yaml:"port" json:"port"`
	TransportProtocol string `yaml:"transport_protocol" json:"transport_protocol"`
}

type Pivio struct {
	ID        string        `yaml:"id" json:"id"`
	ShortName string        `yaml:"short_name" json:"short_name"`
	Runtime   PivioRuntime  `yaml:"runtime" json:"runtime"`
	Services  PivioServices `yaml:"service" json:"service"`
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

	var pivios []Pivio
	if err := json.Unmarshal(data, &pivios); err != nil {
		return nil, err
	}

	return pivios, nil
}
