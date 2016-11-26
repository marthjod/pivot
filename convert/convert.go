package convert

import (
	"fmt"

	"github.com/marthjod/pivot/mapping"
	"github.com/marthjod/pivot/model"
)

func Convert(p *model.Pivio, m *mapping.Mappings) string {
	cpuSize := m.CPU.Values.GetSize(p.Runtime.CPU)
	ramSize := m.RAM.Values.GetSize(p.Runtime.RAM)
	diskSize := m.Disk.Values.GetSize(p.Runtime.Disk)

	// TODO use a template
	ret := fmt.Sprintf("%v: %v, %v: %v, %v: %v",
		m.CPU.Alias, cpuSize,
		m.RAM.Alias, ramSize,
		m.Disk.Alias, diskSize)

	return ret
}
