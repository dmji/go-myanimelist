package mal_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/dmji/go-myanimelist/mal/malhttp"
)

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
	testErrorResponse(t, err, malhttp.ErrorResponse{Message: "manga not found", Err: "not_found"})
}
