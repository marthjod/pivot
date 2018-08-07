package client

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/marthjod/pivot/model"
)

// Client holds information for querying a Pivio API
type Client struct {
	QueryEndpoint string
}

func (c Client) buildShortnameQueryURL(name string) string {
	query := fmt.Sprintf(`{"match": {"short_name": "%s"}}`, name)
	log.Printf("building query `%s`", query)
	return fmt.Sprintf("%s/document?query=%s", c.QueryEndpoint, url.QueryEscape(query))
}

// QueryByShortname queries Pivio API and returns a matching Pivio record(s) on success
func (c Client) QueryByShortname(name string) (model.Pivio, error) {
	url := c.buildShortnameQueryURL(name)

	log.Printf("accessing %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		return model.Pivio{}, err
	}
	defer res.Body.Close()

	// dump, err := httputil.DumpResponse(res, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%q", dump)

	var pivios = []model.Pivio{}
	pivios, err = model.ReadJSONMultiple(res.Body)
	if err != nil || len(pivios) < 1 {
		return model.Pivio{}, err
	}
	return pivios[0], nil
}
