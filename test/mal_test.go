package mal_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal/common"
)

// setup sets up a test HTTP server along with a mal.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup(cls ...*http.Client) (client *mal.Site, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer that the test HTTP server will use
	// to mock API responses.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the MyAnimeList client being tested and is configured to use
	// test server.
	var httpClient *http.Client = nil
	if len(cls) > 0 {
		httpClient = cls[0]
	}

	client = mal.NewSite(httpClient)
	baseURL, _ := url.Parse(server.URL + "/")
	client.SetBaseURL(baseURL)

	return client, mux, server.Close
}

type urlValues map[string]string

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

func testErrorResponse(t *testing.T, err error, want common.ErrorResponse) {
	t.Helper()
	errResp := &common.ErrorResponse{}
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

func testResponseOffset(t *testing.T, resp *common.Response, next, prev int, prefix string) {
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

func testResponseStatusCode(t *testing.T, resp *common.Response, code int, prefix string) {
	t.Helper()
	if resp == nil {
		t.Fatalf("%s resp is nil, want StatusCode=%d", prefix, code)
	}
	if got, want := resp.StatusCode, code; got != want {
		t.Errorf("%s resp.StatusCode=%d, want %d", prefix, got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := mal.NewSite(nil)

	// test default base URL
	if got, want := c.BaseURL(), mal.DefaultBaseURL; got != want {
		t.Errorf("NewClient.BaseURL = %v, want %v", got, want)
	}
}

func TestErrorResponse(t *testing.T) {
	errResp := &common.ErrorResponse{
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
	c := mal.NewSite(nil)

	inURL, outURL := "foo", mal.DefaultBaseURL+"foo"
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
	c := mal.NewSite(nil)
	_, err := c.DirectRequest().NewRequest("invalid method", "/foo")
	if err == nil {
		t.Error("NewRequest with invalid method expected to return err")
	}
}

func TestNewRequestBadEndpoint(t *testing.T) {
	c := mal.NewSite(nil)
	inURL := "%foo"
	_, err := c.DirectRequest().NewRequest("GET", inURL)
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
