package mal_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal_client"
)

type urlValues map[string]string

func mockServer() (url string, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer that the test HTTP server will use
	// to mock API responses.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	baseURL := server.URL + "/"

	return baseURL, mux, server.Close
}

// setup sets up a test HTTP server along with a mal.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup(cls ...*http.Client) (*mal_client.Client, *mal.Site, *http.ServeMux, func()) {
	// client is the MyAnimeList client being tested and is configured to use
	// test server.
	var httpClient *http.Client = nil
	if len(cls) > 0 {
		httpClient = cls[0]
	}

	baseURL, mux, teardown := mockServer()

	reqClient, err := mal_client.NewClientUrl(httpClient, &baseURL)
	if err != nil {
		panic(err)
	}
	malClient, err := mal.NewSite(mal.WithCustomClientPtr(reqClient))
	if err != nil {
		panic(err)
	}

	return reqClient, malClient, mux, teardown
}

func testURLValues(t *testing.T, r *http.Request, values urlValues) {
	t.Helper()
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}
	actual := r.URL.Query()
	if !reflect.DeepEqual(want, actual) {
		t.Errorf("URL Values = %v, want %v", actual, want)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	t.Helper()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request body\nhave: %q\nwant: %q", got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}

func testContentType(t *testing.T, r *http.Request, want string) {
	ct := r.Header.Get("Content-Type")
	if ct != want {
		t.Errorf("Content-Type = %q, want %q", ct, want)
	}
}

func testErrorResponse(t *testing.T, err error, want mal_client.ErrorResponse) {
	t.Helper()
	errResp := &mal_client.ErrorResponse{}
	if !errors.As(err, &errResp) {
		t.Fatalf("err is type %T, want type *ErrorResponse.", err)
	}
	if got, want := errResp.Message, want.Message; got != want {
		t.Errorf("ErrorResponse.Message = %q, want %q", got, want)
	}
	if got, want := errResp.Err, want.Err; got != want {
		t.Errorf("ErrorResponse.Err = %q, want %q", got, want)
	}
}

func testResponseOffset(t *testing.T, resp *mal_client.Response, next, prev int, prefix string) {
	t.Helper()
	if resp == nil {
		t.Fatalf("%s resp is nil, want NextOffset=%d and PrevOffset=%d", prefix, next, prev)
	}
	if got, want := resp.NextOffset, next; got != want {
		t.Errorf("%s resp.NextOffset=%d, want %d", prefix, got, want)
	}
	if got, want := resp.PrevOffset, prev; got != want {
		t.Errorf("%s resp.PrevOffset=%d, want %d", prefix, got, want)
	}
}

func testResponseStatusCode(t *testing.T, resp *mal_client.Response, code int, prefix string) {
	t.Helper()
	if resp == nil {
		t.Fatalf("%s resp is nil, want StatusCode=%d", prefix, code)
	}
	if got, want := resp.StatusCode, code; got != want {
		t.Errorf("%s resp.StatusCode=%d, want %d", prefix, got, want)
	}
}
