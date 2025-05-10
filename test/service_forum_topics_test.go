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

func TestForumServiceTopics(t *testing.T) {
	_, client, mux, teardown := setup()
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
		&mal_prm.ForumTopicsRequestParameters{
			BoardID:       1,
			SubboardID:    1,
			Query:         "foo",
			TopicUserName: "bar",
			UserName:      "baz",
			Limit:         10,
			Sort:          mal_prm.SortTopicsRecent,
		},
	)
	if err != nil {
		t.Errorf("Forum.Topics returned error: %v", err)
	}

	want := []mal_type.Topic{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Forum.Topics returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "Forum.Topics")
}

func TestForumServiceTopicsError(t *testing.T) {
	_, client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forum/topics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, http.StatusInternalServerError)
	})

	ctx := context.Background()
	_, resp, err := client.Forum.Topics(ctx, nil)
	if err == nil {
		t.Fatal("Forum.Topics expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "Forum.Topics")
	testErrorResponse(t, err, mal_client.ErrorResponse{Message: "mal is down", Err: "internal"})
}
