package simple

import (
	"encoding/json"

	"github.com/marthjod/pivot/convert"
	"github.com/marthjod/pivot/model"
)

// Simple represents a simple renderable service definition.
type Simple map[string]map[string]interface{}

// Converter is a converter for simple service definitions.
type Converter struct {
	convert.Converter
	Pivio *model.Pivio
}

// Render implements the converter interface.
func (d Converter) Render() (string, error) {
	return d.convert().toJSON()
}

func (d Converter) convert() Simple {
	return map[string]map[string]interface{}{
		d.Pivio.ShortName: {
			"cpu":    d.Pivio.Runtime.CPU,
			"disk":   d.Pivio.Runtime.Disk,
			"memory": d.Pivio.Runtime.RAM,
			"zone":   d.Pivio.Runtime.NetworkZone,
		},
	}
}

func (d Simple) toJSON() (string, error) {
	y, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(y), nil
}
