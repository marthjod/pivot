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

func (c Converter) Convert(p *model.Pivio) Service {
	zones := []string{
		fmt.Sprintf("zone-%s-%s", strings.ToUpper(p.Runtime.NetworkZone), strings.ToUpper(p.ShortName)),
		"zone-LOGGING",
		"zone-ACCESS",
	}

	sort.Strings(zones)

	cs := map[string]map[string]interface{}{
		fmt.Sprintf("hg-%s", p.ShortName): {
			"cpu": c.Aliases.Get("cpu", p.Runtime.CPU),
			"vcpu": c.Aliases.Get("vcpu", p.Runtime.CPU),
			"image": c.Aliases.Get("image", p.Runtime.Disk),
			"memory": c.Aliases.Get("memory", p.Runtime.RAM),
			"networks": zones,
		},
	}

	return cs
}
