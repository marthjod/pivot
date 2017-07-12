package convert

import (
	"testing"
	"os"
	"log"
)

const (
	aliases  = "../formats/custom/aliases.yaml"
)

var expected = []struct{
	scope string
	key string
	out interface{}
}{
	{
		scope: "doesnotexist",
		key:   "subkey",
		out:   "subkey",
	},
	{
		scope: "memory",
		key: "M",
		out: 4096,
	},
}

func TestGet(t *testing.T) {
	f, err := os.Open(aliases)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()

	a, err := ReadAliases(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, exp := range expected {
		actual := a.Get(exp.scope, exp.key)
		if actual != exp.out {
			t.Errorf("Expected:\n%v\nActual:\n%v", exp.out, actual)
		}
	}
}
