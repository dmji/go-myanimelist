package mal_client_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/dmji/go-myanimelist/mal_client"
)

func TestNewClient(t *testing.T) {
	c, err := mal_client.NewClientUrl(nil, nil)
	if err != nil {
		t.Errorf("Site creation error: %v", err)
		return
	}

	// test default base URL
	if got, want := c.BaseURL(), mal_client.DefaultBaseURL; got != want {
		t.Errorf("NewClient.BaseURL = %v, want %v", got, want)
	}
}

func TestNewClientWringURL(t *testing.T) {
	wrongUrl := "foo\x7fclr"

	_, err := mal_client.NewClientUrl(nil, &wrongUrl)
	if err == nil {
		t.Errorf("Expected creation error for wrong URL: %s", wrongUrl)
		return
	}

	// test default base URL
	want := "parse \"foo\\x7fclr\": net/url: invalid control character in URL"
	got := err.Error()
	if got != want {
		t.Errorf("NewClient.BaseURL = %v, want %v", got, want)
	}
}

func TestErrorResponse(t *testing.T) {
	errResp := &mal_client.ErrorResponse{
		Response: &http.Response{
			Request: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "http",
					Host:   "foo.com",
				},
			},
			StatusCode: 500,
		},
		Message: "service gone",
		Err:     "boom",
	}
	if got, want := errResp.Error(), "GET http://foo.com: 500 service gone boom"; got != want {
		t.Errorf("ErrorResponse.Error() = %q, want %q", got, want)
	}
}
