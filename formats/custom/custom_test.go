package custom

import (
	"github.com/marthjod/pivot/model"
	"os"
	"testing"
	"log"
	"github.com/marthjod/pivot/convert"
)

const (
	pivio    = "../../pivio.yaml"
	aliases  = "aliases.yaml"
)

func TestConvert(t *testing.T) {
	var (
		expected = `service-CFPA:
  cpu: 2
  image: ubuntu-latest-minimal
  memory: 2048
  networks:
  - zone-ACCESS
  - zone-DMZ-CFPA
  - zone-LOGGING
  vcpu: 8
`
	)

	f, err := os.Open(pivio)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()

	p, err := model.Read(f)
	if err != nil {
		t.Fatal(err)
	}

	f, err = os.Open(aliases)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()

	a, err := convert.ReadAliases(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	conv := ServiceConverter{
		Pivio: p,
		Aliases: a,
	}

	actual, err := conv.Render()
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual != expected {
		t.Errorf("Output does not mach expectation. Expected:\n%s\nActual:\n%s", expected, actual)
	}
}
