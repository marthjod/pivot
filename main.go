package main

import (
	"flag"
	"fmt"

	"github.com/marthjod/pivot/convert"
	"github.com/marthjod/pivot/mapping"
	"github.com/marthjod/pivot/model"
)

func main() {
	var (
		err         error
		pivioYAML   string
		mappingYAML string
		pivio       *model.Pivio
		mapp        *mapping.Mappings
	)

	flag.StringVar(&pivioYAML, "pivio", "pivio.yaml", "Path to pivio.yaml")
	flag.StringVar(&mappingYAML, "mapping", "mapping.yaml", "Path to mapping YAML")
	flag.Parse()

	pivio, err = model.FromFile(pivioYAML)
	if err != nil {
		panic(err)
	}

	mapp, err = mapping.FromFile(mappingYAML)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Runtime:\n%+v\nMapping:\n%+v\nConverted:\n%v\n",
		pivio.Runtime, mapp, convert.Convert(pivio, mapp))
}
