package custom

import (
	"fmt"
	"sort"
	"strings"

	"github.com/marthjod/pivot/convert"
	"github.com/marthjod/pivot/model"
	"gopkg.in/yaml.v2"
)

// Service represents a renderable service definition.
type Service map[string]map[string]interface{}

// ServiceConverter knows how to convert a service definition.
type ServiceConverter struct {
	convert.Converter
	Pivio   *model.Pivio
	Aliases *convert.Aliases
}

// Render implements the Converter interface.
func (c *ServiceConverter) Render() (string, error) {
	return c.convert().toYAML()
}

func (c *ServiceConverter) convert() Service {
	zones := []string{
		fmt.Sprintf("zone-%s-%s", strings.ToUpper(c.Pivio.Runtime.NetworkZone), strings.ToUpper(c.Pivio.ShortName)),
		"zone-LOGGING",
		"zone-ACCESS",
	}

	sort.Strings(zones)

	return map[string]map[string]interface{}{
		fmt.Sprintf("service-%s", c.Pivio.ShortName): {
			"cpu":      c.Aliases.Get("cpu", c.Pivio.Runtime.CPU),
			"vcpu":     c.Aliases.Get("vcpu", c.Pivio.Runtime.CPU),
			"image":    c.Aliases.Get("image", c.Pivio.Runtime.Disk),
			"memory":   c.Aliases.Get("memory", c.Pivio.Runtime.RAM),
			"networks": zones,
		},
	}
}

func (s Service) toYAML() (string, error) {
	y, err := yaml.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(y), nil
}
