package convert

import (
	"io"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Aliases map[string]map[string]interface{}

func ReadAliases(r io.Reader) (*Aliases, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	aliases := Aliases{}
	if err := yaml.Unmarshal(data, &aliases); err != nil {
		return nil, err
	}

	return &aliases, nil
}

func (a Aliases) Get(scope, key string) interface{} {
	if val, ok := a[scope][key]; ok {
		return val
	}
	return key
}
