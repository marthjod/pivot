package main

import (
	"flag"
	"github.com/marthjod/pivot/customformat"
	"github.com/marthjod/pivot/model"
	"log"
	"os"
	"text/template"
)

func main() {
	var (
		err          error
		yamlFile     string
		templateFile string
		tpl          *template.Template
		customFormat bool
		pivio        *model.Pivio
	)

	flag.StringVar(&yamlFile, "pivio", "pivio.yaml", "Path to pivio.yaml (input)")
	flag.StringVar(&templateFile, "template", "", "Path to template file for output rendering")
	flag.BoolVar(&customFormat, "custom", false, "Marshal custom YAML output")
	flag.Parse()

	if templateFile == "" && !customFormat {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(yamlFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pivio, err = model.Read(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Read from %s:\n%+v\n", yamlFile, pivio)

	out := ""
	if customFormat {
		customFormat := customformat.Convert(pivio)
		if err != nil {
			log.Fatal(err)
		}
		out, err = customFormat.Yaml()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		tpl, err = template.ParseFiles(templateFile)
		if err != nil {
			log.Fatal(err)
		}

		out, err = pivio.Render(tpl)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Output:\n%s\n", out)
}
