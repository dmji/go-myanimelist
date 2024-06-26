package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal/malhttp"
	"github.com/dmji/go-myanimelist/mal/maltype"
	"github.com/dmji/go-myanimelist/mal/prm"
)

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
	opts := client.Forum.TopicsOptions
	got, resp, err := client.Forum.Topics(ctx,
		opts.BoardID(1),
		opts.SubboardID(1),
		prm.Limit(10),
		prm.Offset(0),
		opts.SortTopics.Recent(),
		opts.Query("foo"),
		opts.TopicUserName("bar"),
		opts.UserName("baz"),
	)
	if err != nil {
		t.Errorf("Forum.Topics returned error: %v", err)
	}

	want := []maltype.Topic{{ID: 1}, {ID: 2}}
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
	testErrorResponse(t, err, malhttp.ErrorResponse{Message: "mal is down", Err: "internal"})
}
