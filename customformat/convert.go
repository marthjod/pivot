package customformat

import (
	"fmt"
	"github.com/marthjod/pivot/model"
	"gopkg.in/yaml.v2"
)

type Service map[string]CustomFormat

type CustomFormat struct {
	CPU      string   `yaml:"cpu"`
	Memory   string   `yaml:"memory"`
	Image    string   `yaml:"image"`
	Networks []string `yaml:"networks"`
}

func (c Service) Yaml() (string, error) {
	y, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(y), nil
}

func Convert(p *model.Pivio) *Service {
	csf := CustomFormat{
		CPU:    p.Runtime.CPU,
		Image:  p.Runtime.Disk,
		Memory: p.Runtime.RAM,
	}
	csf.Networks = append(csf.Networks, p.Runtime.NetworkZone)

	cs := Service{
		fmt.Sprintf("hg-%s", p.ShortName): csf,
	}
	return &cs
}
