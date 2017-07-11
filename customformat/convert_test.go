package customformat

import (
	"github.com/marthjod/pivot/model"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	var (
		yaml     = "../pivio.yaml"
		expected = `hg-CFPA:
  cpu: XL
  memory: S
  image: S
  networks:
  - DMZ
`
	)

	f, err := os.Open(yaml)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()
	p, err := model.Read(f)
	if err != nil {
		t.Fatal(err)
	}

	c := Convert(p)
	actual, err := c.Yaml()
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual != expected {
		t.Errorf("YAML output does not mach expectation. Expected:\n%s\nActual:\n%s", expected, actual)
	}
}
