# go-myanimelist

[![Go Reference](https://pkg.go.dev/badge/github.com/dmji/go-myanimelist/mal.svg)](https://pkg.go.dev/github.com/dmji/go-myanimelist/mal)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmji/go-myanimelist)](https://goreportcard.com/report/github.com/dmji/go-myanimelist)
[![Coverage Status](https://coveralls.io/repos/github/dmji/go-myanimelist/badge.svg?branch=main)](https://coveralls.io/github/dmji/go-myanimelist?branch=main)
[![Test Status](https://github.com/dmji/go-myanimelist/workflows/tests/badge.svg)](https://github.com/dmji/go-myanimelist/actions?query=workflow%3Atests)
[![Integration Status](https://github.com/dmji/go-myanimelist/workflows/integration/badge.svg)](https://github.com/dmji/go-myanimelist/actions?query=workflow%3Aintegration)

go-myanimelist is a Go client library for accessing the [MyAnimeList API v2](https://myanimelist.net/apiconfig/references/api/v2).

## Fork Status

The project files were rebased and the parameters and structures were put into a separate package.
Option providers for service function calls were added.

## Project Status

The project has been updated to support MyAnimeList API v2.

## Installation

This package can be installed using:

    `go get github.com/dmji/go-myanimelist/mal`

## Usage

Import the package using:

```go
import "github.com/dmji/go-myanimelist/mal"
```

First construct a new mal client:

```go
c := mal.NewSite(nil, nil)
```

Then use one of the client's services (User, Anime, Manga and Forum) to access
the different MyAnimeList API methods.

## Authentication

When creating a new client, pass an `http.Client` that can handle authentication
for you.

### Accessing publicly available information

To access public information, you need to add the ` X-MAL-CLIENT-ID` header in
your requests. You can achieve this by creating an `http.Client` with a custom
transport and use it as shown below:

```go
type clientIDTransport struct {
	Transport http.RoundTripper
	ClientID  string
}

func (c *clientIDTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.Transport == nil {
		c.Transport = http.DefaultTransport
	}
	req.Header.Add("X-MAL-CLIENT-ID", c.ClientID)
	return c.Transport.RoundTrip(req)
}

func main() {
	publicInfoClient := &http.Client{
		// Create client ID from https://myanimelist.net/apiconfig.
		Transport: &clientIDTransport{ClientID: "<Your application client ID>"},
	}

	c := mal.NewSite(publicInfoClient, nil)
	// ...
}
```

_(see: example/public_data_only/main.go)_

### Authenticating using OAuth2

The recommended way is to use the `golang.org/x/oauth2` package
(https://github.com/golang/oauth2). After completing the OAuth2 flow, you will
get an oauth2 token containing an access token, a refresh token and an
expiration date. The oauth2 token can easily be stored in JSON format and used
like this:

```go
const storedToken = `
{
	"token_type": "Bearer",
	"access_token": "yourAccessToken",
	"refresh_token": "yourRefreshToken",
	"expiry": "2021-06-01T16:12:56.1319122Z"
}`

oauth2Token := new(oauth2.Token)
_ = json.Unmarshal([]byte(storedToken), oauth2Token)

// Create client ID and secret from https://myanimelist.net/apiconfig.
//
// Secret is currently optional if you choose App Type 'other'.
oauth2Conf := &oauth2.Config{
    ClientID:     "<Enter your registered MyAnimeList.net application client ID>",
    ClientSecret: "<Enter your registered MyAnimeList.net application client secret>",
    Endpoint: oauth2.Endpoint{
        AuthURL:   "https://myanimelist.net/v1/oauth2/authorize",
        TokenURL:  "https://myanimelist.net/v1/oauth2/token",
        AuthStyle: oauth2.AuthStyleInParams,
    },
}

oauth2Client := oauth2Conf.Client(ctx, oauth2Token)

// The oauth2Client will refresh the token if it expires.
c := mal.NewSite(oauth2Client, nil)
```

Note that all calls made by the client above will include the specified oauth2
token which is specific for an authenticated user. Therefore, authenticated
clients should almost never be shared between different users.

Performing the OAuth2 flow involves registering a MAL API application and then
asking for the user's consent to allow the application to access their data.

There is a detailed example of how to perform the Oauth2 flow and get an oauth2
token through the terminal under `example/malauth`. The only thing you need to run
the example is a client ID and a client secret which you can acquire after
registering your MAL API application. Here's how:

1. Navigate to https://myanimelist.net/apiconfig or go to your MyAnimeList
   profile, click Edit Profile and select the API tab on the far right.
2. Click Create ID and submit the form with your application details.

After registering your application, you can run the example and pass the client
ID and client secret through flags:

    cd example/malauth
    go run main.go democlient.go --client-id=... --client-secret=...

    or

    go install github.com/dmji/go-myanimelist/example/malauth
    malauth --client-id=... --client-secret=...

After you perform a successful authentication once, the oauth2 token will be
cached in a file under the same directory which makes it easier to run the
example multiple times.

Official MAL API OAuth2 docs:
https://myanimelist.net/apiconfig/references/authorization

## List

To search and get anime and manga data:

```go
opts := c.Anime.ListOptions
anime, _, err := c.Anime.List(ctx, "hokuto no ken",
	opts.Fields(
		opts.AnimeFields.Rank(),
		opts.AnimeFields.Popularity(),
		opts.AnimeFields.MyListStatus(),
	),
	opts.Limit(5),
)
// ...


opts := c.Manga.ListOptions
anime, _, err := c.Manga.List(ctx, "hokuto no ken",
	opts.Fields(
		opts.MangaFields.Rank(),
		opts.MangaFields.Popularity(),
		opts.MangaFields.MyListStatus(),
	),
	opts.Limit(5),
)
// ...
```

You may get user specific data for a certain record by using the optional field.

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/anime_get
- https://myanimelist.net/apiconfig/references/api/v2#operation/manga_get

## UserList

To get the anime or manga list of a user:

```go
// Get the authenticated user's anime list, filter only watching anime, sort by
// last updated, include list status.
opts := c.User.AnimeListOptions
anime, _, err := c.User.AnimeList(ctx, "@me",
	opts.Fields(opts.UserListFields.ListStatus()),
	opts.AnimeStatus.Watching(),
	opts.SortAnimeList.ByListUpdatedAt(),
	opts.Limit(5),
)
// ...

// Get the authenticated user's manga list's second page, sort by score,
// include list status, comments and tags.
opts := c.User.MangaListOptions
manga, _, err := c.User.MangaList(ctx, "@me",
	opts.SortMangaList.ByListScore(),
	opts.Fields(opts.UserListFields.ListStatus("comments", "tags")),
	opts.Limit(5),
	opts.Offset(1),
)
// ...
```

You may provide the username of the user or "@me" to get the authenticated
user's list.

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_animelist_get
- https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_mangalist_get

## MyInfo

To get information about the authenticated user:

```go
user, _, err := c.User.MyInfo(ctx)
// ...

opts := client.User.MyInfoOptions
user, _, err := client.User.MyInfo(ctx,
	opts.Fields(
		opts.UserFields.TimeZone(),
		opts.UserFields.IsSupporter(),
	),
)
// ...
```

This method can use the Fields option but the API doesn't seem to be able to
send optional fields like "anime_statistics" at the time of writing.

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_get

## Details

To get details for a certain anime or manga:

```go
opts := c.Anime.DetailsOptions
a, _, err := c.Anime.Details(ctx, 967,
	opts.Fields(
		opts.AnimeFields.AlternativeTitles(),
		opts.AnimeFields.MediaType(),
		opts.AnimeFields.NumEpisodes(),
		opts.AnimeFields.StartSeason(),
		opts.AnimeFields.Source(),
		opts.AnimeFields.Genres(),
		opts.AnimeFields.Studios(),
		opts.AnimeFields.AverageEpisodeDuration(),
	),
)
// ...

opts := c.Manga.DetailsOptions
m, _, err := c.Manga.Details(ctx, 401,
	opts.Fields(
		opts.MangaFields.AlternativeTitles(),
		opts.MangaFields.MediaType(),
		opts.MangaFields.NumVolumes(),
		opts.MangaFields.NumChapters(),
		opts.MangaFields.Authors("last_name", "first_name"),
		opts.MangaFields.Genres(),
		opts.MangaFields.Status(),
	),
)
// ...
```

By default most fields are not populated so use the Fields option to request the
fields you need.

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_get
- https://myanimelist.net/apiconfig/references/api/v2#operation/manga_manga_id_get

## Ranking

To get anime or manga based on a certain ranking:

```go
opts := c.Anime.RankingOptions
anime, _, err := c.Anime.Ranking(ctx,
	opts.AnimeRanking.ByPopularity(),
	opts.Fields(
		opts.AnimeFields.Rank(),
		opts.AnimeFields.Popularity(),
	),
	opts.Limit(6),
)
// ...

opts := c.Manga.RankingOptions
manga, _, err := c.Manga.Ranking(ctx,
	opts.MangaRanking.ByPopularity(),
	opts.Fields(
		opts.MangaFields.Rank(),
		opts.MangaFields.Popularity(),
	),
	opts.Limit(6),
)
// ...
```

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/anime_ranking_get
- https://myanimelist.net/apiconfig/references/api/v2#operation/manga_ranking_get

## Add or Update List

To add or update an entry in an authenticated user's list, provide the anime or
manga ID and then options to specify the status, score, comments, tags etc.

```go
opts := c.Anime.UpdateMyListStatusOptions
s, _, err := c.Anime.UpdateMyListStatus(ctx, 967,
	opts.AnimeStatus.Watching(),
	opts.NumEpisodesWatched(73),
	opts.Score(8),
	opts.Comments("You wa shock!"),
	opts.StartDate(time.Date(2022, 02, 20, 0, 0, 0, 0, time.UTC)),
	opts.FinishDate(time.Time{}), // Remove an existing date.
)
// ...

opts := c.Manga.UpdateMyListStatusOptions
s, _, err := c.Manga.UpdateMyListStatus(ctx, 401,
	opts.MangaStatus.Reading(),
	opts.NumVolumesRead(1),
	opts.NumChaptersRead(5),
	opts.Comments("Migi"),
	opts.StartDate(time.Date(2022, 02, 20, 0, 0, 0, 0, time.UTC)),
	opts.FinishDate(time.Time{}), // Remove an existing date.
)
// ...
```

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_my_list_status_put
- https://myanimelist.net/apiconfig/references/api/v2#operation/manga_manga_id_my_list_status_put

## Delete

To delete anime or manga from a user's list, simply provide their IDs:

```go
_, err := c.Anime.DeleteMyListItem(ctx, 967)
// ...

_, err := c.Manga.DeleteMyListItem(ctx, 401)
// ...
```

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_my_list_status_delete
- https://myanimelist.net/apiconfig/references/api/v2#operation/manga_manga_id_my_list_status_delete

## More Examples

See package examples:

- https://pkg.go.dev/github.com/dmji/go-myanimelist/example/

## Unit Testing

To run all unit tests:

`go test .\test\... -cover -coverpkg='./mal/...'`

To see test coverage in your browser:

`go test .\test\... -coverpkg='./mal/...' -covermode=count -coverprofile=count && go tool cover -html count`

## Integration Testing

The integration tests will exercise the entire package against the live
MyAnimeList API. As a result, these tests take much longer to run and there is
also a much higher chance of false positives in test failures due to network
issues etc.

These tests are meant to be run using a dedicated test account that contains
empty anime and manga lists. A valid oauth2 token needs to be provided every
time. Check the authentication section to learn how to get one.

By default the integration tests are skipped when an oauth2 token is not
provided. To run all tests including the integration tests:

` go test --client-id='<your app client ID>``' --oauth2-token='<your oauth2 token>``' `

## License

MIT
