package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_prm"
	"github.com/dmji/go-myanimelist/mal_type"
)

func TestUserServiceAnimeList(t *testing.T) {
	_, client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/foo/animelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"status": "completed",
			"sort":   "anime_id",
			"fields": "id,genres",
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
	got, resp, err := client.User.AnimeList(ctx, "foo",
		&mal_prm.UserAnimeListRequestParameters{
			Fields: []mal_prm.AnimeField{
				mal_prm.AnimeFieldTypeID.AnimeField(),
				mal_prm.AnimeFieldTypeGenres.AnimeField(),
			},
			Limit:  10,
			Offset: 0,
			Sort:   mal_prm.SortAnimeListByAnimeID,
			Status: mal_prm.AnimeStatusCompleted,
			NSFW:   true,
		},
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
	_, client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/foo/animelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, 500)
	})

	ctx := context.Background()
	_, resp, err := client.User.AnimeList(ctx, "foo",
		&mal_prm.UserAnimeListRequestParameters{
			Fields: []mal_prm.AnimeField{
				mal_prm.AnimeFieldTypeID.AnimeField(),
				mal_prm.AnimeFieldTypeGenres.AnimeField(),
			},
			Limit:  10,
			Sort:   mal_prm.SortAnimeListByAnimeID,
			Status: mal_prm.AnimeStatusCompleted,
		},
	)
	if err == nil {
		t.Fatal("User.AnimeList expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "User.AnimeList")
	testErrorResponse(t, err, mal_client.ErrorResponse{Message: "mal is down", Err: "internal"})
}
