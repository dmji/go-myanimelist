package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_type"
)

func TestAnimeServiceDetails(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"fields": "foo,bar",
		})
		testBody(t, r, "")
		fmt.Fprint(w, `{"id":1}`)
	})

	ctx := context.Background()

	a, _, err := client.Anime.Details(ctx, 1,
		client.Anime.DetailsOptions.Fields("foo,bar"),
	)
	if err != nil {
		t.Errorf("Anime.Details returned error: %v", err)
	}
	want := &mal_type.Anime{ID: 1}
	if got := a; !reflect.DeepEqual(got, want) {
		t.Errorf("Anime.Details returned\nhave: %+v\n\nwant: %+v", got, want)
	}
}

func TestAnimeServiceDetailsError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"anime deleted","error":"not_found"}`, 404)
	})

	ctx := context.Background()
	_, _, err := client.Anime.Details(ctx, 1)
	if err == nil {
		t.Fatal("Anime.Details expected not found error, got no error.")
	}
	testErrorResponse(t, err, mal_client.ErrorResponse{Message: "anime deleted", Err: "not_found"})
}
