package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marthjod/pivot/model"
)

// Client holds information for querying a Pivio API
type Client struct {
	QueryEndpoint string
}

func (c Client) buildShortnameQueryURL(name string) string {
	query := "document?&query=%7B%0A%20%20%20%20%22match%22:%20%7B%0A%20%20%20%20%20%20%20%20%22short_name%22:%20%22" + name + "%22%0A%20%20%20%20%7D%0A%7D"
	return fmt.Sprintf("%s/%s", c.QueryEndpoint, query)
}

// QueryByShortname queries Pivio API and returns a matching Pivio record(s) on success
func (c Client) QueryByShortname(name string) ([]model.Pivio, error) {
	url := c.buildShortnameQueryURL(name)

	log.Printf("Accessing %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		return []model.Pivio{}, err
	}

	return model.ReadMultiple(res.Body)
}
