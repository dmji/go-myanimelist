# go-myanimelist

go-myanimelist is a Go client library for accessing the [MyAnimeList API v2](https://myanimelist.net/apiconfig/references/api/v2).

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/nstratos/go-myanimelist/mal?status.svg)](https://godoc.org/github.com/nstratos/go-myanimelist/mal)
[![Go Report Card](https://goreportcard.com/badge/github.com/nstratos/go-myanimelist)](https://goreportcard.com/report/github.com/nstratos/go-myanimelist)
[![Coverage Status](https://coveralls.io/repos/github/nstratos/go-myanimelist/badge.svg?branch=master)](https://coveralls.io/github/nstratos/go-myanimelist?branch=master)
[![Build Status](https://travis-ci.org/nstratos/go-myanimelist.svg?branch=master)](https://travis-ci.org/nstratos/go-myanimelist)

As of March 2017, this package is featured in
[awesome-go](https://github.com/avelino/awesome-go).

## Installation

This package can be installed using:

	go get github.com/nstratos/go-myanimelist/mal

## Usage

Import the package using:

```go
import "github.com/nstratos/go-myanimelist/mal"
```

First construct a new mal client:

```go
c := mal.NewClient(nil)
```

Then use one of the client's services (User, Anime, Manga and Forum) to access
the different MyAnimeList API methods.

## Authentication

When creating a new client, pass an `http.Client` that can handle authentication
for you. The recommended way is to use the `golang.org/x/oauth2` package
(https://github.com/golang/oauth2). After performing the OAuth2 flow, you will
get an access token which can be used like this:

```go
ctx := context.Background()
c := mal.NewClient(
	oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "<your access token>"},
	)),
)
```

Note that all calls made by the client above will include the specified access
token which is specific for an authenticated user. Therefore, authenticated
clients should almost never be shared between different users.

Performing the OAuth2 flow involves registering a MAL API application and then
asking for the user's consent to allow the application to access their data.

There is a detailed example of how to perform the Oauth2 flow and get an access
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

    go install github.com/nstratos/go-myanimelist/example/malauth
    malauth --client-id=... --client-secret=...

After you perform a successful authentication once, the access token will be
cached in a file under the same directory which makes it easier to run the
example multiple times.

Official MAL API OAuth2 docs:
https://myanimelist.net/apiconfig/references/authorization

## List

To search and get anime and manga data:

```go
list, _, err := c.Anime.List(ctx, "hokuto no ken",
	mal.Fields{"rank", "popularity", "my_list_status"},
	mal.Limit(5),
)
// ...

list, _, err := c.Manga.List(ctx, "hokuto no ken",
	mal.Fields{"rank", "popularity", "my_list_status"},
	mal.Limit(5),
)
// ...
```

You may get user specific data for a certain record by using the optional field
"my_list_status".

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/anime_get
- https://myanimelist.net/apiconfig/references/api/v2#operation/manga_get

## Details

To get details for a certain anime or manga:

```go
a, _, err := c.Anime.Details(ctx, 967,
	mal.Fields{
		"alternative_titles",
		"media_type",
		"num_episodes",
		"start_season",
		"source",
		"genres",
		"studios",
		"average_episode_duration",
	},
)
// ...

m, _, err := c.Manga.Details(ctx, 401,
	mal.Fields{
		"alternative_titles",
		"media_type",
		"num_volumes",
		"num_chapters",
		"authors{last_name, first_name}",
		"genres",
		"status",
	},
)
// ...
```

By default most fields are not populated so use the Fields option to request the
fields you need.

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_get
- https://myanimelist.net/apiconfig/references/api/v2#operation/manga_manga_id_get

## Add

To add anime and manga, you provide their IDs and values through AnimeEntry and
MangaEntry:

```go
c := mal.NewClient(mal.Auth("<your username>", "<your password>"))

_, err := c.Anime.Add(9989, mal.AnimeEntry{Status: mal.Current, Episode: 1})
// ...

_, err := c.Manga.Add(35733, mal.MangaEntry{Status: mal.Planned, Chapter: 1, Volume: 1})
// ...
```

Note that when adding entries, Status is required.

## Update

Similar to Add, Update also needs the ID of the entry and the values to be
updated:

```go
c := mal.NewClient(mal.Auth("<your username>", "<your password>"))

_, err := c.Anime.Update(9989, mal.AnimeEntry{Status: mal.Completed, Score: 9})
// ...

_, err := c.Manga.Update(35733, mal.MangaEntry{Status: mal.OnHold})
// ...
```

## Delete

To delete anime and manga, simply provide their IDs:

```go
c := mal.NewClient(mal.Auth("<your username>", "<your password>"))

_, err := c.Anime.Delete(9989)
// ...

_, err := c.Manga.Delete(35733)
// ...
```

## More Examples

See package examples:
https://godoc.org/github.com/nstratos/go-myanimelist/mal#pkg-examples

## Advanced Control

If you need more control over the created requests, you can use an option to
pass a custom HTTP client to NewClient:

```go
c := mal.NewClient(&http.Client{})
```

For example this http.Client will make sure to cancel any request that takes
longer than 1 second:

```go
httpcl := &http.Client{
	Timeout: 1 * time.Second,
}
c := mal.NewClient(httpcl)
// ...
```

## Unit Testing

To run all unit tests:

	go test -cover

To see test coverage in your browser:

	go test -covermode=count -coverprofile=count.out && go tool cover -html count.out

## Integration Testing

The integration tests will exercise the entire package against the live
MyAnimeList API. As a result, these tests take much longer to run and there is
also a much higher chance of false positives in test failures due to network
issues etc.

These tests are meant to be run using a dedicated test account that contains
empty anime and manga lists. A valid access token needs to be provided every
time.

To run the integration tests:

	go test --access-token '<your access token>'

## License

MIT
