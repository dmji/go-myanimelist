package mal_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/dmji/go-myanimelist/mal/malhttp"
)

func TestAnimeServiceDeleteMyListItem(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/1/my_list_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	ctx := context.Background()
	resp, err := client.Anime.DeleteMyListItem(ctx, 1)
	if err != nil {
		t.Errorf("Anime.DeleteMyListItem returned error: %v", err)
	}
	testResponseStatusCode(t, resp, http.StatusOK, "Anime.DeleteMyListItem")
}

func TestAnimeServiceDeleteMyListItemError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/anime/1/my_list_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		http.Error(w, `{"message":"anime not found","error":"not_found"}`, http.StatusNotFound)
	})

	ctx := context.Background()
	resp, err := client.Anime.DeleteMyListItem(ctx, 1)
	if err == nil {
		t.Fatal("Anime.DeleteMyListItem expected internal error, got no error.")
	}
	testResponseStatusCode(t, resp, http.StatusNotFound, "Anime.DeleteMyListItem")
	testErrorResponse(t, err, malhttp.ErrorResponse{Message: "anime not found", Err: "not_found"})
}
