package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal/api_driver"
	"github.com/dmji/go-myanimelist/mal/containers"
)

func TestMangaServiceList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"q":      "query",
			"fields": "foo,bar",
			"limit":  "10",
			"offset": "0",
			"nsfw":   "true",
		})
		const out = `
		{
		  "data": [
		    {
		      "node": { "id": 1 }
		    },
		    {
		      "node": { "id": 2 }
		    }
		  ],
		  "paging": {
		    "next": "?offset=4",
		    "previous": "?offset=2"
		  }
		}`
		fmt.Fprint(w, out)
	})

	ctx := context.Background()
	opts := client.Manga.ListOptions
	got, resp, err := client.Manga.List(ctx, "query",
		opts.Fields("foo", "bar"),
		opts.Limit(10),
		opts.Offset(0),
		opts.NSFW(true),
	)
	if err != nil {
		t.Errorf("Manga.List returned error: %v", err)
	}
	want := []containers.Manga{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Manga.List returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "Manga.List")
}

func TestMangaServiceListError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, 500)
	})

	ctx := context.Background()
	_, resp, err := client.Manga.List(ctx, "query")
	if err == nil {
		t.Fatal("Manga.List expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "Manga.List")
	testErrorResponse(t, err, api_driver.ErrorResponse{Message: "mal is down", Err: "internal"})
}
