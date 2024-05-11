package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal/maltype"
)

func TestUserServiceMangaList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/foo/mangalist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"status": "completed",
			"sort":   "manga_id",
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
			    "status": "plan_to_read"
			  }
		    },
		    {
		      "node": { "id": 2 },
			  "list_status": {
			    "status": "reading"
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
	opts := client.User.MangaListOptions
	got, resp, err := client.User.MangaList(ctx, "foo",
		opts.MangaStatus.Completed(),
		opts.SortMangaList.ByMangaID(),
		opts.Fields("foo", "bar"),
		opts.Limit(10),
		opts.Offset(0),
		opts.NSFW(true),
	)
	if err != nil {
		t.Errorf("User.MangaList returned error: %v", err)
	}
	want := []maltype.UserManga{
		{
			Manga:  maltype.Manga{ID: 1},
			Status: maltype.MangaListStatus{Status: "plan_to_read"},
		},
		{
			Manga:  maltype.Manga{ID: 2},
			Status: maltype.MangaListStatus{Status: "reading"},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("User.MangaList returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "User.MangaList")
}
