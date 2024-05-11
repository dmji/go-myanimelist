package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal/maltype"
)

func TestAnimeServiceSeasonal(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/season/2020/summer", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"sort":   "anime_num_list_users",
			"fields": "foo,bar",
			"limit":  "10",
			"offset": "0",
			"nsfw":   "false",
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
		    "next": "?offset=4"
		  },
		  "season": {
			"year": 2020,
			"season": "summer"
		  }
		}`
		fmt.Fprint(w, out)
	})

	ctx := context.Background()
	opts := client.Anime.SeasonalOptions
	got, resp, err := client.Anime.Seasonal(ctx, 2020, opts.AnimeSeason.Summer(),
		opts.SortSeasonalAnime.ByUsersCount(),
		opts.Fields("foo", "bar"),
		opts.Limit(10),
		opts.Offset(0),
		opts.NSFW(false),
	)
	if err != nil {
		t.Errorf("Anime.Seasonal returned error: %v", err)
	}
	want := []maltype.Anime{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Anime.Seasonal returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 0, "Anime.Seasonal")
}
