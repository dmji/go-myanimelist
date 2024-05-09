package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal/common"
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
	want := &mal.Forum{
		Categories: []mal.ForumCategory{
			{
				Title: "MyAnimeList",
				Boards: []mal.ForumBoard{
					{
						ID:          17,
						Title:       "MAL Guidelines",
						Description: "Site rules.",
						Subboards:   []mal.ForumSubboard{{ID: 2, Title: "Anime DB"}},
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
	testErrorResponse(t, err, common.ErrorResponse{Message: "forum deleted", Err: "not_found"})
}

func TestForumServiceTopicDetails(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forum/topic/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"limit":  "10",
			"offset": "0",
		})
		const out = `
		{
		  "data": {
		    "title": "Best posts",
		    "posts": [{ "id": 1 }, { "id": 2 }]
		  },
		  "paging": {
		    "next": "?offset=4",
			"previous": "?offset=2"
		  }
		}`
		fmt.Fprint(w, out)
	})

	ctx := context.Background()
	got, resp, err := client.Forum.TopicDetails(ctx, 1,
		common.Limit(10),
		common.Offset(0),
	)
	if err != nil {
		t.Errorf("Forum.TopicDetails returned error: %v", err)
	}
	want := mal.TopicDetails{
		Title: "Best posts",
		Posts: []mal.Post{{ID: 1}, {ID: 2}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Forum.TopicDetails returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "Forum.TopicDetails")
}

func TestForumServiceTopicDetailsError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forum/topic/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, http.StatusInternalServerError)
	})

	ctx := context.Background()
	_, resp, err := client.Forum.TopicDetails(ctx, 1)
	if err == nil {
		t.Fatal("Forum.TopicDetails expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "Forum.TopicDetails")
	testErrorResponse(t, err, common.ErrorResponse{Message: "mal is down", Err: "internal"})
}

func TestForumServiceTopics(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forum/topics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"board_id":        "1",
			"subboard_id":     "1",
			"limit":           "10",
			"offset":          "0",
			"sort":            "recent",
			"q":               "foo",
			"topic_user_name": "bar",
			"user_name":       "baz",
		})
		const out = `
		{
		  "data": [{ "id": 1 }, { "id": 2 }],
		  "paging": {
		    "next": "?offset=4",
			"previous": "?offset=2"
		  }
		}`
		fmt.Fprint(w, out)
	})

	ctx := context.Background()
	got, resp, err := client.Forum.Topics(ctx,
		mal.BoardID(1),
		mal.SubboardID(1),
		common.Limit(10),
		common.Offset(0),
		mal.SortTopicsRecent,
		mal.Query("foo"),
		mal.TopicUserName("bar"),
		mal.UserName("baz"),
	)
	if err != nil {
		t.Errorf("Forum.Topics returned error: %v", err)
	}
	want := []mal.Topic{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Forum.Topics returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "Forum.Topics")
}

func TestForumServiceTopicsError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forum/topics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, http.StatusInternalServerError)
	})

	ctx := context.Background()
	_, resp, err := client.Forum.Topics(ctx)
	if err == nil {
		t.Fatal("Forum.Topics expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "Forum.Topics")
	testErrorResponse(t, err, common.ErrorResponse{Message: "mal is down", Err: "internal"})
}
