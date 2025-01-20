package mal_client_test

import (
	"io"
	"net/url"
	"testing"

	"github.com/dmji/go-myanimelist/mal_client"
)

func TestNewRequest(t *testing.T) {
	c, err := mal_client.NewClientUrl(nil, nil)
	if err != nil {
		t.Errorf("Site creation error: %v", err)
		return
	}

	inURL, outURL := "foo", mal_client.DefaultBaseURL+"foo"
	inBody, outBody := func(v *url.Values) { v.Set("name", "bar") }, "name=bar"

	req, err := c.NewRequest("GET", inURL, inBody)
	if err != nil {
		t.Fatalf("NewRequest(%q) returned error: %v", inURL, err)
	}

	// test that the endpoint URL was correctly added to the base URL
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL = %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest("+`func(v *url.Values) { v.Set("name", "bar")`+") Body \nhave: %q\nwant: %q", got, want)
	}

	// test that Content-Type header is correctly set when body is set
	if got, want := req.Header.Get("Content-Type"), "application/x-www-form-urlencoded"; got != want {
		t.Errorf("NewRequest() Content-Type \nhave: %q\nwant: %q", got, want)
	}
}

func TestNewRequestInvalidMethod(t *testing.T) {
	c, err := mal_client.NewClientUrl(nil, nil)
	if err != nil {
		t.Errorf("Site creation error: %v", err)
		return
	}

	_, err = c.NewRequest("invalid method", "/foo")
	if err == nil {
		t.Error("NewRequest with invalid method expected to return err")
	}
}

func TestNewRequestBadEndpoint(t *testing.T) {
	c, err := mal_client.NewClientUrl(nil, nil)
	if err != nil {
		t.Errorf("Site creation error: %v", err)
		return
	}

	inURL := "%foo"
	_, err = c.NewRequest("GET", inURL)
	if err == nil {
		t.Errorf("NewRequest(%q) should return parse err", inURL)
	}
}
