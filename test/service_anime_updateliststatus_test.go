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
)

func TestAnimeServiceUpdateMyListStatus(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/1/my_list_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testContentType(t, r, "application/x-www-form-urlencoded")
		testBody(t, r, "comments=comments&finish_date=&is_rewatching=true&num_times_rewatched=2&num_watched_episodes=3&priority=2&rewatch_value=1&score=8&start_date=2022-02-20&status=completed&tags=foo%2Cbar")
		const out = `
		{
		  "status": "completed",
		  "score": 8,
		  "num_episodes_watched": 3,
		  "is_rewatching": true,
		  "updated_at": "2018-04-25T15:59:52Z",
		  "start_date": "2022-02-20",
		  "priority": 2,
		  "num_times_rewatched": 2,
		  "rewatch_value": 1,
		  "tags": ["foo","bar"],
		  "comments": "comments"
		}`
		fmt.Fprint(w, out)
	})

	ctx := context.Background()
	opts := client.Anime.UpdateMyListStatusOptions
	got, _, err := client.Anime.UpdateMyListStatus(ctx, 1,
		opts.AnimeStatus.Completed(),
		opts.IsRewatching(true),
		opts.Score(8),
		opts.NumEpisodesWatched(3),
		opts.Priority(2),
		opts.NumTimesRewatched(2),
		opts.RewatchValue.VeryLow(),
		opts.Tags("foo", "bar"),
		opts.Comments("comments"),
		opts.StartDate(time.Date(2022, 02, 20, 0, 0, 0, 0, time.UTC)),
		opts.FinishDate(time.Time{}),
	)
	if err != nil {
		t.Errorf("Anime.UpdateMyListStatus returned error: %v", err)
	}

	want := &containers.AnimeListStatus{
		Status:             opts.AnimeStatus.Completed(),
		IsRewatching:       true,
		Score:              8,
		NumEpisodesWatched: 3,
		Priority:           2,
		NumTimesRewatched:  2,
		RewatchValue:       1,
		Tags:               []string{"foo", "bar"},
		Comments:           "comments",
		UpdatedAt:          time.Date(2018, 04, 25, 15, 59, 52, 0, time.UTC),
		StartDate:          "2022-02-20",
		FinishDate:         "",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Anime.UpdateMyListStatus returned\nhave: %+v\n\nwant: %+v", got, want)
	}
}

func TestAnimeServiceUpdateMyListStatusError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/1/my_list_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		http.Error(w, `{"message":"mal is down","error":"internal"}`, 500)
	})

	ctx := context.Background()
	_, resp, err := client.Anime.UpdateMyListStatus(ctx, 1)
	if err == nil {
		t.Fatal("Anime.UpdateMyListStatus expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusInternalServerError, "Anime.UpdateMyListStatus")
	testErrorResponse(t, err, api_driver.ErrorResponse{Message: "mal is down", Err: "internal"})
}
