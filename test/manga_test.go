package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal/api_driver"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

func TestMangaServiceDetails(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"fields": "foo,bar",
		})
		testBody(t, r, "")
		fmt.Fprint(w, `{"id":1}`)
	})

	ctx := context.Background()
	a, _, err := client.Manga.Details(ctx, 1, prm.Fields{"foo,bar"})
	if err != nil {
		t.Errorf("Manga.Details returned error: %v", err)
	}
	want := &containers.Manga{ID: 1}
	if got := a; !reflect.DeepEqual(got, want) {
		t.Errorf("Manga.Details returned\nhave: %+v\n\nwant: %+v", got, want)
	}
}

func TestMangaServiceDetailsError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"manga deleted","error":"not_found"}`, 404)
	})

	ctx := context.Background()
	_, _, err := client.Manga.Details(ctx, 1)
	if err == nil {
		t.Fatal("Manga.Details expected not found error, got no error.")
	}
	testErrorResponse(t, err, api_driver.ErrorResponse{Message: "manga deleted", Err: "not_found"})
}

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
	got, resp, err := client.Manga.List(ctx, "query",
		prm.Fields{"foo", "bar"},
		prm.Limit(10),
		prm.Offset(0),
		prm.NSFW(true),
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

func TestMangaServiceRanking(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/ranking", func(w http.ResponseWriter, r *http.Request) {
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
	got, resp, err := client.Manga.Ranking(ctx, prm.MangaRankingAll,
		prm.Fields{"foo", "bar"},
		prm.Limit(10),
		prm.Offset(0),
	)
	if err != nil {
		t.Errorf("Manga.Ranking returned error: %v", err)
	}
	want := []containers.Manga{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Manga.Ranking returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 0, "Manga.Ranking")
}
