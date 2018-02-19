package main

import (
	"flag"
	"log"

	"github.com/marthjod/pivot/client"
)

func main() {
	var (
		queryEndpoint = flag.String("a", "http://pivio.example.com:9123", "Pivio query API endpoint")
		shortname     = flag.String("s", "AAA", "Artifact shortname to query")
	)
	flag.Parse()

	c := client.Client{QueryEndpoint: *queryEndpoint}
	pivio, err := c.QueryByShortname(*shortname)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", pivio)
}
