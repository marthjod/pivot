package mapping

import (
	"io/ioutil"

	"strings"

	"gopkg.in/yaml.v2"
)

type TShirtSizes map[string]int

func GetSize(t TShirtSizes, key string) int {
	if val, found := t[key]; found {
		return val
	} else if val, found = t[strings.ToLower(key)]; found {
		return val
	} else {
		return t[strings.ToUpper(key)]
	}
}

type TShirtSizeMapping struct {
	Alias  string      `yaml:"rename_key_to"`
	Values TShirtSizes `yaml:"values"`
}

type CPUMapping struct {
	Alias       string      `yaml:"rename_key_to"`
	Values      TShirtSizes `yaml:"values"`
	RatioFactor float32     `yaml:"ratio_factor"`
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
