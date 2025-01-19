//go:build testuse

package mal_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal_client"
)

func TestNewClient(t *testing.T) {
	c, err := mal.NewSite(nil, nil)
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
	_, err := mal.NewSite(nil, &wrongUrl)
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

func TestNewRequest(t *testing.T) {
	c, err := mal.NewSite(nil, nil)
	if err != nil {
		t.Errorf("Site creation error: %v", err)
		return
	}

	inURL, outURL := "foo", mal_client.DefaultBaseURL+"foo"
	inBody, outBody := func(v *url.Values) { v.Set("name", "bar") }, "name=bar"

	req, err := c.DirectRequest().NewRequest("GET", inURL, inBody)
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
	c, err := mal.NewSite(nil, nil)
	if err != nil {
		t.Errorf("Site creation error: %v", err)
		return
	}
	_, err = c.DirectRequest().NewRequest("invalid method", "/foo")
	if err == nil {
		t.Error("NewRequest with invalid method expected to return err")
	}
}

func TestNewRequestBadEndpoint(t *testing.T) {
	c, err := mal.NewSite(nil, nil)
	if err != nil {
		t.Errorf("Site creation error: %v", err)
		return
	}

	inURL := "%foo"
	_, err = c.DirectRequest().NewRequest("GET", inURL)
	if err == nil {
		t.Errorf("NewRequest(%q) should return parse err", inURL)
	}
}

func TestDo(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	type foo struct {
		Bar string `json:"bar"`
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if want := "GET"; r.Method != want {
			t.Errorf("request method = %v, want %v", r.Method, want)
		}
		fmt.Fprint(w, `{"bar":"&bull; foobar"}`)
	})

	req, _ := client.DirectRequest().NewRequest("GET", "/")

	body := new(foo)
	ctx := context.Background()
	_, err := client.DirectRequest().Do(ctx, req, body)
	if err != nil {
		t.Fatalf("Do() returned err = %v", err)
	}

	want := &foo{"&bull; foobar"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Do() response body = %v, want %v", body, want)
	}
}

func TestDoHTTPError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad request", http.StatusBadRequest)
	})

	req, _ := client.DirectRequest().NewRequest("GET", "/")

	ctx := context.Background()
	resp, err := client.DirectRequest().Do(ctx, req, nil)
	if err == nil {
		t.Fatal("Expected HTTP 400 error, got no error.")
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected HTTP 400 error, got %d status code.", resp.StatusCode)
	}
}

type errTransport struct{}

func (e errTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("connection refused")
}

func TestDoRoundTripError(t *testing.T) {
	client, mux, teardown := setup(
		&http.Client{
			Transport: &errTransport{},
		},
	)
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	req, _ := client.DirectRequest().NewRequest("GET", "/")
	ctx := context.Background()
	_, err := client.DirectRequest().Do(ctx, req, nil)
	if err == nil {
		t.Error("Expected connection refused error.")
	}
}

func TestDoNoContent(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.DirectRequest().NewRequest("GET", ".")
	ctx := context.Background()
	_, err := client.DirectRequest().Do(ctx, req, &body)
	if err != nil {
		t.Fatalf("Do returned unexpected error: %v", err)
	}
}

func TestDoDecodeError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "this is not JSON")
	})

	var body json.RawMessage

	req, _ := client.DirectRequest().NewRequest("GET", ".")
	ctx := context.Background()
	_, err := client.DirectRequest().Do(ctx, req, &body)
	if err == nil {
		t.Fatal("Expected JSON decode error.")
	}
}

func TestDoNilContext(t *testing.T) {
	client, _, teardown := setup()
	defer teardown()

	req, _ := client.DirectRequest().NewRequest("GET", ".")
	var ctx context.Context = nil
	_, err := client.DirectRequest().Do(ctx, req, nil)
	if err == nil {
		t.Errorf("Do should return error when we pass nil context.")
	}
}
