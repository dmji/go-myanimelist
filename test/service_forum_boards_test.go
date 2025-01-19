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

func TestForumServiceBoards(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forum/boards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{})
		testBody(t, r, "")
		const out = `
		{
		  "categories": [
		    {
		  	  "title": "MyAnimeList",
		  	  "boards": [
		        {
		          "id": 17,
		          "title": "MAL Guidelines",
		          "description": "Site rules.",
		          "subboards": [{"id": 2,"title": "Anime DB"}]
		  	    }
		      ]
		    }
		  ]
		}`
		fmt.Fprint(w, out)
	})

	ctx := context.Background()
	a, _, err := client.Forum.Boards(ctx)
	if err != nil {
		t.Errorf("Forum.Boards returned error: %v", err)
	}
	want := &mal_type.Forum{
		Categories: []mal_type.ForumCategory{
			{
				Title: "MyAnimeList",
				Boards: []mal_type.ForumBoard{
					{
						ID:          17,
						Title:       "MAL Guidelines",
						Description: "Site rules.",
						Subboards:   []mal_type.ForumSubboard{{ID: 2, Title: "Anime DB"}},
					},
				},
			},
		},
	}
	if got := a; !reflect.DeepEqual(got, want) {
		t.Errorf("Forum.Boards returned\nhave: %+v\n\nwant: %+v", got, want)
	}
}

func TestForumServiceBoardsError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forum/boards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"forum deleted","error":"not_found"}`, 404)
	})

	ctx := context.Background()
	_, _, err := client.Forum.Boards(ctx)
	if err == nil {
		t.Fatal("Forum.Boards expected not found error, got no error.")
	}
	testErrorResponse(t, err, mal_client.ErrorResponse{Message: "forum deleted", Err: "not_found"})
}
