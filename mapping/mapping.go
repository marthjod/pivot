package mapping

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type TShirtSizes struct {
	S int `yaml:"s"`
	M int `yaml:"m"`
	L int `yaml:"l"`
	// TODO get custom
}

func (t *TShirtSizes) GetSize(alias string) int {
	switch alias {
	case "S":
		return t.S
	case "M":
		return t.M
	case "L":
		return t.L
	default:
		return 0
	}
}

type TShirtSizeMapping struct {
	Alias  string      `yaml:"rename_key_to"`
	Values TShirtSizes `yaml:"values"`
}

type CPUMapping struct {
	TShirtSizeMapping
	RatioFactor float32 `yaml:"ratio_factor"`
}

type Mappings struct {
	RAM  TShirtSizeMapping `yaml:"ram"`
	CPU  CPUMapping        `yaml:"cpu"`
	Disk TShirtSizeMapping `yaml:"disk"`
}

type MappingInfo struct {
	Mappings Mappings `yaml:"mapping"`
}

func FromFile(path string) (*Mappings, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	mappingInfo := MappingInfo{}
	if err := yaml.Unmarshal(data, &mappingInfo); err != nil {
		return nil, err
	}

	return &mappingInfo.Mappings, nil
}
