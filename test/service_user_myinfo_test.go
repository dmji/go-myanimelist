package mal_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_type"
)

func TestUserServiceMyInfo(t *testing.T) {
	_, client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/@me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURLValues(t, r, urlValues{
			"fields": "time_zone,is_supporter",
		})
		fmt.Fprint(w, `{"id":1}`)
	})

	ctx := context.Background()
	opts := client.User.MyInfoOptions
	u, _, err := client.User.MyInfo(ctx,
		opts.Fields(
			opts.UserFields.TimeZone(),
			opts.UserFields.IsSupporter(),
		),
	)
	if err != nil {
		t.Errorf("User.MyInfo returned error: %v", err)
	}
	want := &mal_type.User{ID: 1}
	if got := u; !reflect.DeepEqual(got, want) {
		t.Errorf("User.MyInfo returned\nhave: %+v\n\nwant: %+v", got, want)
	}
}

func TestUserServiceMyInfoError(t *testing.T) {
	_, client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/@me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"message":"","error":"not_found"}`, 404)
	})

	ctx := context.Background()
	_, _, err := client.User.MyInfo(ctx)
	if err == nil {
		t.Fatal("User.MyInfo expected not found error, got no error.")
	}
	testErrorResponse(t, err, mal_client.ErrorResponse{Err: "not_found"})
}
