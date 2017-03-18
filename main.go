package main

import (
	"flag"
	"github.com/marthjod/pivot/model"
	"log"
	"text/template"
)

func main() {
	var (
		err          error
		pivioYAML    string
		templateFile string
		tpl          *template.Template
		pivio        *model.Pivio
	)

	flag.StringVar(&pivioYAML, "pivio", "pivio.yaml", "Path to pivio.yaml")
	flag.StringVar(&templateFile, "template", "templates/network-dsl.tpl", "Path to output template")
	flag.Parse()

	log.Println("reading from " + pivioYAML)
	pivio, err = model.PivioFromFile(pivioYAML)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%q\n", pivio.ShortName)
	log.Printf("  provides %v", pivio.Services.Provides)
	log.Printf("  internal deps: %v\n", pivio.Services.DependsOn.Internal)
	log.Printf("  external deps: %v\n", pivio.Services.DependsOn.External)

	tpl, err = template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	rendered, err := pivio.Render(tpl)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rendered)
}
