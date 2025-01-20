package mal_test

import (
	"context"
	"fmt"

	"github.com/dmji/go-myanimelist/mal"
)

func ExampleSite_User_myinfo() {
	ctx := context.Background()

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()

	c, err := mal.NewSite(mal.WithCustomClientUrl(nil, &server.URL))
	if err != nil {
		fmt.Printf("Site creation error: %v", err)
		return
	}

	user, _, err := c.User.MyInfo(ctx)
	if err != nil {
		fmt.Printf("User.MyInfo error: %v", err)
		return
	}
	fmt.Printf("ID: %5d, Joined: %v, Username: %s\n", user.ID, user.JoinedAt.Format("Jan 2006"), user.Name)
	// Output:
	// ID: 4592783, Joined: May 2015, Username: nstratos
}
