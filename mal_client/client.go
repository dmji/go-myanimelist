package mal_client

import (
	"net/http"
	"net/url"

	"github.com/dmji/qs"
)

const (
	// DefaultBaseURL is the default MyAnimeList API base URL.
	DefaultBaseURL = "https://api.myanimelist.net/v2/"
)

// Client manages communication with the MyAnimeList API.
type Client struct {
	client       *http.Client
	baseURL      *url.URL
	urlMarshaler *qs.QSMarshaler
}

// BaseURL returns the base url of the http.client active request url. By default, this is reference to server MyAnimeList API
func (c *Client) BaseURL() string {
	return c.baseURL.String()
}

// NewClient returns a new MyAnimeList API client. The httpClient parameter
func NewClient(httpClient *http.Client, baseURL *url.URL) *Client {
	urlMarshaler := qs.NewMarshaler(&qs.MarshalOptions{},
		qs.WithMarshalOptionSliceSeparator(qs.OptionSliceSeparatorComma),
	)

	return &Client{
		client:       httpClient,
		baseURL:      baseURL,
		urlMarshaler: urlMarshaler,
	}
}

func NewClientUrl(httpClient *http.Client, baseURL *string) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	if baseURL == nil {
		defaultURL := DefaultBaseURL
		baseURL = &defaultURL
	}

	baseRelURL, err := url.Parse(*baseURL)
	if err != nil {
		return nil, err
	}

	return NewClient(httpClient, baseRelURL), nil
}
