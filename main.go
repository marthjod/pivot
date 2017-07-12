package main

import (
	"flag"
	"github.com/marthjod/pivot/formats/custom"
	"github.com/marthjod/pivot/formats/simple"
	"github.com/marthjod/pivot/model"
	"log"
	"os"
	"github.com/marthjod/pivot/convert"
)

func main() {
	var (
		err         error
		pivioYaml   = flag.String("pivio", "pivio.yaml", "Path to pivio.yaml (input)")
		aliasesYaml = flag.String("aliases", "aliases.yaml", "Path to alias mapping (input)")
		outputFormat = flag.String("format", "default", "Conversion output format")
		pivio       *model.Pivio
		converter convert.Converter
	)

	flag.Parse()

	f, err := os.Open(*pivioYaml)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pivio, err = model.Read(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Read from %s:\n%+v\n", *pivioYaml, pivio)
	switch *outputFormat {
	case "custom":
		f, err = os.Open(os.ExpandEnv(*aliasesYaml))
		if err != nil {
			log.Fatal(err.Error())
		}
		aliases, err := convert.ReadAliases(f)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("Read from %s:\n%+v\n", *aliasesYaml, aliases)

		converter = &custom.ServiceConverter{
			Pivio: pivio,
			Aliases: aliases,
		}
	default:
		converter = &simple.Converter{
			Pivio: pivio,
		}
	}

	out, err := converter.Render()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Output:\n%s\n", out)
}
