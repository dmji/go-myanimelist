package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_type"
)

func TestAnimeServiceList(t *testing.T) {
	_, client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime", func(w http.ResponseWriter, r *http.Request) {
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
	opts := client.Anime.ListOptions
	got, resp, err := client.Anime.List(ctx, "query",
		opts.Fields("foo", "bar"),
		opts.Limit(10),
		opts.Offset(0),
		opts.NSFW(true),
	)
	if err != nil {
		t.Errorf("Anime.List returned error: %v", err)
	}
	want := []mal_type.Anime{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Anime.List returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "Anime.List")
}

func TestAnimeServiceListParsePagingError(t *testing.T) {
	tests := []struct {
		name string
		out  string
	}{
		{
			name: "cannot parse next url",
			out: `{
			  "data": [],
			  "paging": { "next": "\f", "previous": "?offset=2" }
			}`,
		},
		{
			name: "cannot parse previous url",
			out: `{
			  "data": [],
			  "paging": { "next": "?offset=2", "previous": "\f" }
			}`,
		},
		{
			name: "cannot parse next offset as int",
			out: `{
			  "data": [],
			  "paging": { "next": "?offset=foo", "previous": "?offset=2" }
			}`,
		},
		{
			name: "cannot parse previous offset as int",
			out: `{
			  "data": [],
			  "paging": { "next": "?offset=2", "previous": "?offset=foo" }
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, client, mux, teardown := setup()
			defer teardown()

			mux.HandleFunc("/anime", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				fmt.Fprint(w, tt.out)
			})

			ctx := context.Background()
			_, _, err := client.Anime.List(ctx, "query")
			if err == nil {
				t.Fatal("Anime.List expected paging error, got no error.")
			}
			if wantPrefix := "paging:"; !strings.HasPrefix(err.Error(), wantPrefix) {
				t.Errorf("Anime.List expected error to start with %q, error is %q", wantPrefix, err.Error())
			}
		})
	}
}

func TestAnimeServiceListError(t *testing.T) {
	_, client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, 500)
	})

	ctx := context.Background()
	_, resp, err := client.Anime.List(ctx, "query")
	if err == nil {
		t.Fatal("Anime.List expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "Anime.List")
	testErrorResponse(t, err, mal_client.ErrorResponse{Message: "mal is down", Err: "internal"})
}
