package malhttp

import (
	"fmt"
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

// Paging provides access to the next and previous page URLs when there are
// pages of results.
type paging struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// MARK: Error format

// An ErrorResponse reports an error caused by an API request.
//
// https://myanimelist.net/apiconfig/references/api/v2#section/Common-formats - Error format
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"`
	Err      string         `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Err)
}

// Response wraps http.Response and is returned in all the library functions
// that communicate with the MyAnimeList API. Even if an error occurs the
// response will always be returned along with the actual error so that the
// caller can further inspect it if needed. For the same reason it also keeps
// a copy of the http.Response.Body that was read when the response was first
// received.
type Response struct {
	*http.Response
	Body []byte

	NextOffset int
	PrevOffset int
}
