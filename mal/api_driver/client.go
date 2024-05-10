package api_driver

import (
	"net/http"
	"net/url"
)

const (
	DefaultBaseURL = "https://api.myanimelist.net/v2/"
)

type Client struct {
	client *http.Client

	// Base URL for MyAnimeList API requests.
	BaseURL *url.URL
}

func NewClient(httpClient *http.Client, baseUrl *url.URL) *Client {
	return &Client{
		client:  httpClient,
		BaseURL: baseUrl,
	}
}

func fillValues(v *url.Values, urlOptions ...func(*url.Values)) {
	for _, o := range urlOptions {
		o(v)
	}
}
