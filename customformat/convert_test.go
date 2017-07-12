package customformat

import (
	"github.com/marthjod/pivot/model"
	"os"
	"testing"
	"log"
)

const (
	pivio    = "../pivio.yaml"
	aliases  = "../customformat/aliases.yaml"
)

func TestConvert(t *testing.T) {
	var (
		expected = `hg-CFPA:
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

	a, err := Read(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	conv := Converter{
		Pivio: p,
		Aliases: a,
	}
	c := conv.Convert(p)
	actual, err := c.Yaml()
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual != expected {
		t.Errorf("YAML output does not mach expectation. Expected:\n%s\nActual:\n%s", expected, actual)
	}
}
