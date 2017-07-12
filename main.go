package main

import (
	"flag"
	"github.com/marthjod/pivot/customformat"
	"github.com/marthjod/pivot/model"
	"log"
	"os"
)

func main() {
	var (
		err         error
		pivioYaml   = flag.String("pivio", "pivio.yaml", "Path to pivio.yaml (input)")
		aliasesYaml = flag.String("aliases", "aliases.yaml", "Path to alias mapping (input)")
		pivio       *model.Pivio
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

	f, err = os.Open(*aliasesYaml)
	if err != nil {
		log.Fatal(err.Error())
	}
	aliases, err := customformat.Read(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Read from %s:\n%+v\n", *aliasesYaml, aliases)

	converter := customformat.Converter{
		Pivio:   pivio,
		Aliases: aliases,
	}
	custom := converter.Convert(pivio)
	if err != nil {
		log.Fatal(err)
	}

	out, err := custom.Yaml()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Output:\n%s\n", out)
}
