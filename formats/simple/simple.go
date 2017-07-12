package simple

import (
	"fmt"
	"github.com/marthjod/pivot/model"
	"github.com/marthjod/pivot/convert"
	"encoding/json"
)

type Simple map[string]map[string]interface{}

func (d Simple) Json() (string, error) {
	y, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(y), nil
}

type Converter struct {
	convert.Converter
	Pivio *model.Pivio
}

func (d Converter) Render() (string, error) {
	return d.Convert().Json()
}

func (d Converter) Convert() Simple {
	return map[string]map[string]interface{}{
		fmt.Sprintf("hg-%s", d.Pivio.ShortName): {
			"cpu": d.Pivio.Runtime.CPU,
			"disk": d.Pivio.Runtime.Disk,
			"memory": d.Pivio.Runtime.RAM,
			"zone": d.Pivio.Runtime.NetworkZone,
		},
	}
}
