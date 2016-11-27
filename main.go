package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/marthjod/pivot/mapping"
	"github.com/marthjod/pivot/model"
)

func main() {
	var (
		err           error
		pivioYAML     string
		mappingYAML   string
		convertedYAML string
		pivio         *model.Pivio
		mapp          *mapping.Mappings
		oneTemplate   *model.OpenNebulaTemplate
		oneYAML       []byte
	)

	flag.StringVar(&pivioYAML, "pivio", "pivio.yaml", "Path to pivio.yaml")
	flag.StringVar(&mappingYAML, "mapping", "mapping.yaml", "Path to mapping YAML")
	flag.StringVar(&convertedYAML, "converted", "converted.yaml", "Path to conversion output")
	flag.Parse()

	pivio, err = model.PivioFromFile(pivioYAML)
	if err != nil {
		panic(err)
	}

	mapp, err = mapping.FromFile(mappingYAML)
	if err != nil {
		panic(err)
	}

	oneTemplate = pivio.Convert(mapp)
	oneYAML, err = oneTemplate.ToYAML()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Runtime:\n%+v\nMapping:\n%+v\nConverted:\n%s\n",
		pivio.Runtime, mapp, oneYAML)

	fmt.Printf("Writing converted YAML to %s\n", convertedYAML)
	err = ioutil.WriteFile(convertedYAML, oneYAML, 0666)
	if err != nil {
		panic(err)
	}

}
