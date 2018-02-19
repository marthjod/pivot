package simple

import (
	"os"
	"testing"

	"github.com/marthjod/pivot/model"
)

const (
	pivio = "../../examples/filebackend/pivio.yaml"
)

func TestConvert(t *testing.T) {
	var (
		expected = `{
  "CFPA": {
    "cpu": "XL",
    "disk": "S",
    "memory": "S",
    "zone": "DMZ"
  }
}`
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

	conv := Converter{
		Pivio: p,
	}

	actual, err := conv.Render()
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual != expected {
		t.Errorf("Output does not mach expectation. Expected:\n%s\nActual:\n%s", expected, actual)
	}
}
