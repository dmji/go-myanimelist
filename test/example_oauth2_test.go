package mal_test

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dmji/go-myanimelist/mal"
	"golang.org/x/oauth2"
)

func newOAuth2Client(ctx context.Context) *http.Client {
	// In order to create a client ID and secret for your application:
	//
	//  1. Navigate to https://myanimelist.net/apiconfig or go to your MyAnimeList
	//     profile, click Edit Profile and select the API tab on the far right.
	//  2. Click Create ID and submit the form with your application details.
	malSecret := clientID
	if clientSecret != nil {
		malSecret = clientSecret
	}

	conf := &oauth2.Config{
		// "<Enter your MyAnimeList.net application client ID>" (now load from argiment)
		ClientID: *clientID,
		// "<Enter your MyAnimeList.net application client secret>" (now load from argiment)
		ClientSecret: *malSecret, // Optional if you chose App Type 'other'.
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://myanimelist.net/v1/oauth2/authorize",
			TokenURL:  "https://myanimelist.net/v1/oauth2/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	// To get your first token you need to complete the oauth2 flow. There is a
	// detailed example that uses the terminal under `example/malauth` which you
	// should adjust for your application.
	//
	// Here we assume we have already received our first valid token and stored
	// it somewhere in JSON format.
	// Comes from secret storage
	//const oauth2Token = `
	//{
	//	"token_type": "Bearer",
	//	"access_token": "yourAccessToken",
	//	"refresh_token": "yourRefreshToken",
	//	"expiry": "2021-06-01T16:12:56.1319122Z"
	//}`

	// Decode stored token to oauth2.Token struct.
	token := new(oauth2.Token)
	_ = json.Unmarshal([]byte(*oauth2Token), token)

	// The oauth2 client returned here with the above configuration and valid
	// token will refresh the token seamlessly when it expires.
	return conf.Client(ctx, token)
}

func Example_oAuth2() {
	ctx := context.Background()
	oauth2Client := newOAuth2Client(ctx)

	c := mal.NewSite(oauth2Client)

	user, _, err := c.User.MyInfo(ctx)
	if err != nil {
		fmt.Printf("User.MyInfo error: %v", err)
		return
	}
	fmt.Printf("ID: %5d, Joined: %v, Username: %s\n", user.ID, user.JoinedAt.Format("Jan 2006"), user.Name)
	// Output:
	// ID: 18315605, Joined: May 2024, Username: go_api_test
}
