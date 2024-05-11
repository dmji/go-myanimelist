package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dmji/go-myanimelist/mal"
)

type clientIDTransport struct {
	Transport http.RoundTripper
	ClientID  string
}

func (c *clientIDTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.Transport == nil {
		c.Transport = http.DefaultTransport
	}
	req.Header.Add("X-MAL-CLIENT-ID", c.ClientID)
	return c.Transport.RoundTrip(req)
}

func main() {
	publicInfoClient := &http.Client{
		// Create client ID from https://myanimelist.net/apiconfig.
		Transport: &clientIDTransport{ClientID: "<Your application client ID>"},
	}

	c, err := mal.NewSite(publicInfoClient, nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	p := c.User.MyInfoOptions
	user, _, err := c.User.MyInfo(ctx, p.Fields(p.ID()))
	if err != nil {
		fmt.Printf("User is not available without autorization. Error: %v\n", err)
	} else {
		fmt.Printf("User ID: %d\n", user.ID)
	}

	li := c.Anime.DetailsOptions
	desc, _, err := c.Anime.Details(ctx, 2019,
		li.Fields(
			li.AnimeFields.Title(),
			li.AnimeFields.Genres(),
		),
	)
	if err != nil {
		fmt.Printf("Anime is not available. Error: %v\n", err)
	} else {
		fmt.Printf("Anime title: %s\n", desc.Title)
		for i, g := range desc.Genres {
			fmt.Printf("Anime #%d genre: %s\n", i, g.Name)
		}
	}
}
