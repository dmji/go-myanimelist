package api_driver

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If data
// is passed as an argument then it will first be encoded in XML and then added
// to the request body as URL encoded value data=<xml>...
// This is how the MyAnimeList requires to receive the data when adding or
// updating entries.
func (c *Client) NewRequest(method, urlStr string, urlOptions ...func(v *url.Values)) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var r io.Reader
	if len(urlOptions) != 0 {
		v := &url.Values{}
		fillValues(v, urlOptions...)
		r = strings.NewReader(v.Encode())
	}

	req, err := http.NewRequest(method, u.String(), r)
	if err != nil {
		return nil, err
	}

	if len(urlOptions) != 0 {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return req, nil
}
