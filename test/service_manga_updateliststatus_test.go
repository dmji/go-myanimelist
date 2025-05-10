package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_type"
)

func TestUserServiceMangaListError(t *testing.T) {
	_, client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/foo/mangalist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, 500)
	})

	ctx := context.Background()
	_, resp, err := client.User.MangaList(ctx, "foo", nil)
	if err == nil {
		t.Fatal("User.MangaList expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "User.MangaList")
	testErrorResponse(t, err, mal_client.ErrorResponse{Message: "mal is down", Err: "internal"})
}

func TestMangaServiceUpdateMyListStatus(t *testing.T) {
	_, client, mux, teardown := setup()
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
	opts := client.Manga.UpdateMyListStatusOptions
	got, _, err := client.Manga.UpdateMyListStatus(ctx, 1,
		opts.MangaStatus.Completed(),
		opts.IsRereading(true),
		opts.Score(8),
		opts.NumVolumesRead(3),
		opts.NumChaptersRead(3),
		opts.Priority(2),
		opts.NumTimesReread(2),
		opts.RereadValue.VeryLow(),
		opts.Tags("foo", "bar"),
		opts.Comments("comments"),
		opts.StartDate(time.Date(2022, 0o2, 20, 0, 0, 0, 0, time.UTC)),
		opts.FinishDate(time.Time{}),
	)
	if err != nil {
		t.Errorf("Manga.UpdateMyListStatus returned error: %v", err)
	}

	want := &mal_type.MangaListStatus{
		Status:          opts.MangaStatus.Completed(),
		IsRereading:     true,
		Score:           8,
		NumVolumesRead:  3,
		NumChaptersRead: 3,
		Priority:        2,
		NumTimesReread:  2,
		RereadValue:     1,
		Tags:            []string{"foo", "bar"},
		Comments:        "comments",
		UpdatedAt:       time.Date(2018, 0o4, 25, 15, 59, 52, 0, time.UTC),
		StartDate:       "2022-02-20",
		FinishDate:      "",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Manga.UpdateMyListStatus returned\nhave: %+v\n\nwant: %+v", got, want)
	}
}

func TestMangaServiceUpdateMyListStatusError(t *testing.T) {
	_, client, mux, teardown := setup()
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
	testErrorResponse(t, err, mal_client.ErrorResponse{Message: "mal is down", Err: "internal"})
}
