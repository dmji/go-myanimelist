package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/dmji/go-myanimelist/mal/api_driver"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
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
	got, resp, err := client.User.MangaList(ctx, "foo",
		prm.MangaStatusCompleted,
		prm.SortMangaListByMangaID,
		prm.Fields{"foo", "bar"},
		prm.Limit(10),
		prm.Offset(0),
		prm.NSFW(true),
	)
	if err != nil {
		t.Errorf("User.MangaList returned error: %v", err)
	}
	want := []containers.UserManga{
		{
			Manga:  containers.Manga{ID: 1},
			Status: containers.MangaListStatus{Status: "plan_to_read"},
		},
		{
			Manga:  containers.Manga{ID: 2},
			Status: containers.MangaListStatus{Status: "reading"},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("User.MangaList returned\nhave: %+v\n\nwant: %+v", got, want)
	}
	testResponseOffset(t, resp, 4, 2, "User.MangaList")
}

func TestUserServiceMangaListError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/foo/mangalist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, 500)
	})

	ctx := context.Background()
	_, resp, err := client.User.MangaList(ctx, "foo")
	if err == nil {
		t.Fatal("User.MangaList expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "User.MangaList")
	testErrorResponse(t, err, api_driver.ErrorResponse{Message: "mal is down", Err: "internal"})
}
func TestMangaServiceUpdateMyListStatus(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/1/my_list_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testContentType(t, r, "application/x-www-form-urlencoded")
		testBody(t, r, "comments=comments&finish_date=&is_rereading=true&num_chapters_read=3&num_times_reread=2&num_volumes_read=3&priority=2&reread_value=1&score=8&start_date=2022-02-20&status=completed&tags=foo%2Cbar")
		const out = `
		{
		  "status": "completed",
		  "score": 8,
		  "num_volumes_read": 3,
		  "num_chapters_read": 3,
		  "is_rereading": true,
		  "updated_at": "2018-04-25T15:59:52Z",
		  "start_date": "2022-02-20",
		  "priority": 2,
		  "num_times_reread": 2,
		  "reread_value": 1,
		  "tags": ["foo","bar"],
		  "comments": "comments"
		}`
		fmt.Fprint(w, out)
	})

	ctx := context.Background()
	got, _, err := client.Manga.UpdateMyListStatus(ctx, 1,
		prm.MangaStatusCompleted,
		prm.IsRereading(true),
		prm.Score(8),
		prm.NumVolumesRead(3),
		prm.NumChaptersRead(3),
		prm.Priority(2),
		prm.NumTimesReread(2),
		prm.RereadValue(1),
		prm.Tags{"foo", "bar"},
		prm.Comments("comments"),
		prm.StartDate(time.Date(2022, 02, 20, 0, 0, 0, 0, time.UTC)),
		prm.FinishDate(time.Time{}),
	)
	if err != nil {
		t.Errorf("Manga.UpdateMyListStatus returned error: %v", err)
	}

	want := &containers.MangaListStatus{
		Status:          prm.MangaStatusCompleted,
		IsRereading:     true,
		Score:           8,
		NumVolumesRead:  3,
		NumChaptersRead: 3,
		Priority:        2,
		NumTimesReread:  2,
		RereadValue:     1,
		Tags:            []string{"foo", "bar"},
		Comments:        "comments",
		UpdatedAt:       time.Date(2018, 04, 25, 15, 59, 52, 0, time.UTC),
		StartDate:       "2022-02-20",
		FinishDate:      "",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Manga.UpdateMyListStatus returned\nhave: %+v\n\nwant: %+v", got, want)
	}
}

func TestMangaServiceUpdateMyListStatusError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/1/my_list_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, 500)
	})

	ctx := context.Background()
	_, resp, err := client.Manga.UpdateMyListStatus(ctx, 1)
	if err == nil {
		t.Fatal("Manga.UpdateMyListStatus expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "Manga.UpdateMyListStatus")
	testErrorResponse(t, err, api_driver.ErrorResponse{Message: "mal is down", Err: "internal"})
}

func TestMangaServiceDeleteMyListItem(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/1/my_list_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	ctx := context.Background()
	resp, err := client.Manga.DeleteMyListItem(ctx, 1)
	if err != nil {
		t.Errorf("Manga.DeleteMyListItem returned error: %v", err)
	}
	testResponseStatusCode(t, resp, http.StatusOK, "Manga.DeleteMyListItem")
}

func TestMangaServiceDeleteMyListItemError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/1/my_list_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		http.Error(w, `{"message":"manga not found","error":"not_found"}`, http.StatusNotFound)
	})

	ctx := context.Background()
	resp, err := client.Manga.DeleteMyListItem(ctx, 1)
	if err == nil {
		t.Fatal("Manga.DeleteMyListItem expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusNotFound, "Manga.DeleteMyListItem")
	testErrorResponse(t, err, api_driver.ErrorResponse{Message: "manga not found", Err: "not_found"})
}
