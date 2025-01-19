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

func TestUserServiceAnimeList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/foo/animelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"status": "completed",
			"sort":   "anime_id",
			"fields": "foo,bar",
			"limit":  "10",
			"offset": "0",
			"nsfw":   "true",
		})
		const out = `
		{
		  "data": [
		    {
		      "node": { "id": 1 },
			  "list_status": {
			    "status": "plan_to_watch"
			  }
		    },
		    {
		      "node": { "id": 2 },
			  "list_status": {
			    "status": "watching"
			  }
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
	opts := client.User.AnimeListOptions
	got, resp, err := client.User.AnimeList(ctx, "foo",
		opts.AnimeStatus.Completed(),
		opts.SortAnimeList.ByAnimeID(),
		opts.Fields("foo", "bar"),
		opts.Limit(10),
		opts.Offset(0),
		opts.NSFW(true),
	)
	if err != nil {
		t.Errorf("User.AnimeList returned error: %v", err)
	}
	want := []mal_type.UserAnime{
		{
			Anime:  mal_type.Anime{ID: 1},
			Status: mal_type.AnimeListStatus{Status: "plan_to_watch"},
		},
		{
			Anime:  mal_type.Anime{ID: 2},
			Status: mal_type.AnimeListStatus{Status: "watching"},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("User.AnimeList returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "User.AnimeList")
}

func TestUserServiceAnimeListError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/foo/animelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, 500)
	})

	ctx := context.Background()
	opts := client.User.AnimeListOptions
	_, resp, err := client.User.AnimeList(ctx, "foo",
		opts.AnimeStatus.Completed(),
		opts.SortAnimeList.ByAnimeID(),
		opts.Fields("foo", "bar"),
		opts.Limit(10),
		opts.Offset(0),
	)
	if err == nil {
		t.Fatal("User.AnimeList expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "User.AnimeList")
	testErrorResponse(t, err, mal_client.ErrorResponse{Message: "mal is down", Err: "internal"})
}
