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

func TestMangaServiceDetails(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"fields": "foo,bar",
		})
		testBody(t, r, "")
		fmt.Fprint(w, `{"id":1}`)
	})

	ctx := context.Background()

	a, _, err := client.Manga.Details(ctx, 1, client.Manga.DetailsOptions.Fields("foo", "bar"))
	if err != nil {
		t.Errorf("Manga.Details returned error: %v", err)
	}
	want := &maltype.Manga{ID: 1}
	if got := a; !reflect.DeepEqual(got, want) {
		t.Errorf("Manga.Details returned\nhave: %+v\n\nwant: %+v", got, want)
	}
}

func TestMangaServiceDetailsError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/manga/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"manga deleted","error":"not_found"}`, 404)
	})

	ctx := context.Background()
	_, _, err := client.Manga.Details(ctx, 1)
	if err == nil {
		t.Fatal("Manga.Details expected not found error, got no error.")
	}
	testErrorResponse(t, err, malhttp.ErrorResponse{Message: "manga deleted", Err: "not_found"})
}
