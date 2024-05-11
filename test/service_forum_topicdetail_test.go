package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal/malhttp"
	"github.com/dmji/go-myanimelist/mal/maltype"
)

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
	opts := client.Forum.TopicDetailsOptions
	got, resp, err := client.Forum.TopicDetails(ctx, 1,
		opts.Limit(10),
		opts.Offset(0),
	)
	if err != nil {
		t.Errorf("Forum.TopicDetails returned error: %v", err)
	}
	want := maltype.TopicDetails{
		Title: "Best posts",
		Posts: []maltype.Post{{ID: 1}, {ID: 2}},
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
	testErrorResponse(t, err, malhttp.ErrorResponse{Message: "mal is down", Err: "internal"})
}
