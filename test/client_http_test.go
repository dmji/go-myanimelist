package mal_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	client, _, mux, teardown := setup()
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

	req, _ := client.NewRequest("GET", "/")

	body := new(foo)
	ctx := context.Background()
	_, err := client.Do(ctx, req, body)
	if err != nil {
		t.Fatalf("Do() returned err = %v", err)
	}

	want := &foo{"&bull; foobar"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Do() response body = %v, want %v", body, want)
	}
}

func TestDoHTTPError(t *testing.T) {
	client, _, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad request", http.StatusBadRequest)
	})

	req, _ := client.NewRequest("GET", "/")

	ctx := context.Background()
	resp, err := client.Do(ctx, req, nil)
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
	client, _, mux, teardown := setup(
		&http.Client{
			Transport: &errTransport{},
		},
	)
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	req, _ := client.NewRequest("GET", "/")
	ctx := context.Background()
	_, err := client.Do(ctx, req, nil)
	if err == nil {
		t.Error("Expected connection refused error.")
	}
}

func TestDoNoContent(t *testing.T) {
	client, _, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", ".")
	ctx := context.Background()
	_, err := client.Do(ctx, req, &body)
	if err != nil {
		t.Fatalf("Do returned unexpected error: %v", err)
	}
}

func TestDoDecodeError(t *testing.T) {
	client, _, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "this is not JSON")
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", ".")
	ctx := context.Background()
	_, err := client.Do(ctx, req, &body)
	if err == nil {
		t.Fatal("Expected JSON decode error.")
	}
}

func TestDoNilContext(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	req, _ := client.NewRequest("GET", ".")
	var ctx context.Context = nil
	_, err := client.Do(ctx, req, nil)
	if err == nil {
		t.Errorf("Do should return error when we pass nil context.")
	}
}
