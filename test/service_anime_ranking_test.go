package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal/maltype"
)

func TestAnimeServiceRanking(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/ranking", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"ranking_type": "all",
			"fields":       "foo,bar",
			"limit":        "10",
			"offset":       "0",
		})
		const out = `
		{
		  "data": [
		    {
		      "node": { "id": 1 },
			  "ranking": { "rank": 1 }
		    },
		    {
		      "node": { "id": 2 },
			  "ranking": { "rank": 2 }
		    }
		  ],
		  "paging": {
		    "next": "?offset=4"
		  }
		}`
		fmt.Fprint(w, out)
	})

	ctx := context.Background()
	opts := client.Anime.RankingOptions
	got, resp, err := client.Anime.Ranking(ctx, opts.AnimeRanking.All(),
		opts.Fields("foo", "bar"),
		opts.Limit(10),
		opts.Offset(0),
	)
	if err != nil {
		t.Errorf("Anime.Ranking returned error: %v", err)
	}
	want := []maltype.Anime{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Anime.Ranking returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 0, "Anime.Ranking")
}
