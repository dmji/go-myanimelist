go-myanimelist
==============

go-myanimelist is a Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/nstratos/go-myanimelist/mal?status.svg)](https://godoc.org/github.com/nstratos/go-myanimelist/mal)
[![Go Report Card](https://goreportcard.com/badge/github.com/nstratos/go-myanimelist)](https://goreportcard.com/report/github.com/nstratos/go-myanimelist)
[![Coverage Status](https://coveralls.io/repos/github/nstratos/go-myanimelist/badge.svg?branch=master)](https://coveralls.io/github/nstratos/go-myanimelist?branch=master)
[![Build Status](https://travis-ci.org/nstratos/go-myanimelist.svg?branch=master)](https://travis-ci.org/nstratos/go-myanimelist)

Installation 
------------

This package can be installed using:

	go get github.com/nstratos/go-myanimelist/mal

Usage
-----

Import the package using:

```go
import "github.com/nstratos/go-myanimelist/mal"
```

First construct a new mal client:

```go
c := mal.NewClient(nil)
```

Then use one of the client's services (Account, Anime or Manga) to access the
different MyAnimeList API methods.

For example, to get the anime and manga list of the user "Xinil":

```go
c := mal.NewClient(nil)

list, _, err := c.Anime.List("Xinil")
// ...

list, _, err := c.Manga.List("Xinil")
// ...
```

If a method requires authentication, make sure to set your MyAnimeList username
and password on the client.

For example to search for anime and manga (needs authentication):

```go
c := mal.NewClient(nil)
c.SetCredentials("<your username>", "<your password>")

result, _, err := c.Anime.Search("bebop")
// ...

result, _, err := c.Manga.Search("bebop")
// ...
```

For more complex searches, you can provide the % operator which is escaped as
%% in Go. Note: This is an undocumented API feature.

```go
c := mal.NewClient(nil)
c.SetCredentials("<your username>", "<your password>")

result, _, err := c.Anime.Search("fate%%heaven%%flower")
// ...
// Will return: Fate/stay night Movie: Heaven's Feel - I. presage flower
```

If you need more control, when creating a new client you can pass an
http.Client as an argument.

For example this http.Client passed to the mal client will make sure to cancel
any request that takes longer than 1 second:

```go
httpcl := &http.Client{
	Timeout: 1 * time.Second,
}
c := mal.NewClient(httpcl)
// ...
```

See more examples: https://godoc.org/github.com/nstratos/go-myanimelist/mal#pkg-examples

Unit Testing
------------

To run all unit tests:

	cd $GOPATH/src/github.com/nstratos/go-myanimelist/mal
	go test -cover

To see test coverage in your browser:

	go test -covermode=count -coverprofile=count.out && go tool cover -html count.out

Integration Testing
-------------------

The integration tests will exercise the entire package against the live
MyAnimeList API. As a result, these tests take much longer to run and there is
also a much higher chance of false positives in test failures due to network
issues etc.

These tests are meant to be run using a dedicated test account that contains
empty anime and manga lists. The username and password of the test account need
to be provided every time.

To run the integration tests:

	cd $GOPATH/src/github.com/nstratos/go-myanimelist/mal
	go test -tags=integration -username '<test account username>' -password '<test account password>'

License
-------

MIT
