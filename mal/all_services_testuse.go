//go:build testuse

package mal

import (
	"context"
	"net/http"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_client"
)

// HTTPDriver is the interface that wraps the Do and NewRequest methods.
//
// Hack to http.client, not recommend to use
type HTTPDriver interface {
	Do(ctx context.Context, req *http.Request, v interface{}) (*mal_client.Response, error)
	NewRequest(method, urlStr string, urlOptions ...func(v *url.Values)) (*http.Request, error)
}

// DirectRequest returns the underlying http.client interface
func (c *Site) DirectRequest() HTTPDriver {
	return c.client
}

// BaseURL returns the base url of the http.client active request url. By default, this is reference to server MyAnimeList API
func (c *Site) BaseURL() string {
	return c.client.BaseURL.String()
}
