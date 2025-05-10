package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal_prm"
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
		Transport: &clientIDTransport{ClientID: "b9cbb74688f2c32bed5e0864c6b5d0b3"},
	}

	c, err := mal.NewSite(mal.WithCustomClientUrl(publicInfoClient, nil))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	user, _, err := c.User.MyInfo(ctx, &mal_prm.UserMyInfoRequestParameters{
		Fields: []mal_prm.UserField{
			{
				Field: mal_prm.UserFieldTypeID,
			},
		},
	})
	if err != nil {
		fmt.Printf("User is not available without autorization. Error: %v\n", err)
	} else {
		fmt.Printf("User ID: %d\n", user.ID)
	}

	desc, _, err := c.Anime.Details(ctx, 2019,
		&mal_prm.AnimeDetailsRequestParameters{
			Fields: []mal_prm.AnimeField{
				mal_prm.AnimeFieldTypeTitle.AnimeField(),
				mal_prm.AnimeFieldTypeGenres.AnimeField(),
			},
		},
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
