package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal/maltype"
)

func TestAnimeServiceSuggested(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/suggestions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"fields": "foo,bar",
			"limit":  "10",
			"offset": "0",
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
	opts := client.Anime.SuggestedOptions
	got, resp, err := client.Anime.Suggested(ctx,
		opts.Fields("foo", "bar"),
		opts.Limit(10),
		opts.Offset(0),
	)
	if err != nil {
		t.Errorf("Anime.Suggested returned error: %v", err)
	}
	want := []maltype.Anime{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Anime.Suggested returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "Anime.Suggested")
}
