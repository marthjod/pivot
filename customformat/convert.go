package customformat

import (
	"fmt"
	"github.com/marthjod/pivot/model"
	"gopkg.in/yaml.v2"
	"strings"
	"sort"
)

type Service map[string]map[string]interface{}

type Converter struct {
	Pivio *model.Pivio
	Aliases *Aliases
}

func (s Service) Yaml() (string, error) {
	y, err := yaml.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(y), nil
}

func (c Converter) Convert() Service {
	zones := []string{
		fmt.Sprintf("zone-%s-%s", strings.ToUpper(c.Pivio.Runtime.NetworkZone), strings.ToUpper(c.Pivio.ShortName)),
		"zone-LOGGING",
		"zone-ACCESS",
	}

	sort.Strings(zones)

	cs := map[string]map[string]interface{}{
		fmt.Sprintf("hg-%s", c.Pivio.ShortName): {
			"cpu": c.Aliases.Get("cpu", c.Pivio.Runtime.CPU),
			"vcpu": c.Aliases.Get("vcpu", c.Pivio.Runtime.CPU),
			"image": c.Aliases.Get("image", c.Pivio.Runtime.Disk),
			"memory": c.Aliases.Get("memory", c.Pivio.Runtime.RAM),
			"networks": zones,
		},
	}

	return cs
}
