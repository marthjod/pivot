package custom

import (
	"fmt"
	"github.com/marthjod/pivot/model"
	"github.com/marthjod/pivot/convert"
	"gopkg.in/yaml.v2"
	"strings"
	"sort"
)

type Service map[string]map[string]interface{}

func (s Service) Yaml() (string, error) {
	y, err := yaml.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(y), nil
}

type ServiceConverter struct {
	convert.Converter
	Pivio *model.Pivio
	Aliases *convert.Aliases
}

func (c *ServiceConverter) Render() (string, error) {
	return c.Convert().Yaml()
}

func (c *ServiceConverter) Convert() Service {
	zones := []string{
		fmt.Sprintf("zone-%s-%s", strings.ToUpper(c.Pivio.Runtime.NetworkZone), strings.ToUpper(c.Pivio.ShortName)),
		"zone-LOGGING",
		"zone-ACCESS",
	}

	sort.Strings(zones)

	return map[string]map[string]interface{}{
		fmt.Sprintf("service-%s", c.Pivio.ShortName): {
			"cpu": c.Aliases.Get("cpu", c.Pivio.Runtime.CPU),
			"vcpu": c.Aliases.Get("vcpu", c.Pivio.Runtime.CPU),
			"image": c.Aliases.Get("image", c.Pivio.Runtime.Disk),
			"memory": c.Aliases.Get("memory", c.Pivio.Runtime.RAM),
			"networks": zones,
		},
	}
}
