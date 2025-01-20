package mal

import (
	"net/http"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_client"
)

type initOptions struct {
	c client
}

func (o *initOptions) initEmptyFields() error {
	var err error
	if o.c == nil {
		o.c, err = mal_client.NewClientUrl(nil, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

type fnOptionApply func(*initOptions) error

func WithCustomClientPtr(c client) fnOptionApply {
	return func(o *initOptions) error {
		o.c = c
		return nil
	}
}

func WithCustomClientUrl(httpClient *http.Client, baseURL *string) fnOptionApply {
	return func(o *initOptions) error {
		var err error
		o.c, err = mal_client.NewClientUrl(httpClient, baseURL)
		return err
	}
}

func WithCustomClient(httpClient *http.Client, baseURL *url.URL) fnOptionApply {
	return func(o *initOptions) error {
		o.c = mal_client.NewClient(httpClient, baseURL)
		return nil
	}
}
