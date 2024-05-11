package mal_test

import (
	"context"
	"fmt"
	"time"

	"github.com/dmji/go-myanimelist/mal"
)

func ExampleSite_User_animelist() {
	ctx := context.Background()

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()

	c, err := mal.NewSite(nil, &server.URL)
	if err != nil {
		fmt.Printf("Site creation error: %v", err)
		return
	}

	opts := c.User.AnimeListOptions
	anime, _, err := c.User.AnimeList(ctx, "@me",
		opts.Fields(opts.UserListFields.ListStatus()),
		opts.SortAnimeList.ByListUpdatedAt(),
		opts.Limit(5),
	)
	if err != nil {
		fmt.Printf("User.AnimeList error: %v", err)
		return
	}
	for _, a := range anime {
		fmt.Printf("ID: %5d, Status: %15q, Episodes Watched: %3d %s\n", a.Anime.ID, a.Status.Status, a.Status.NumEpisodesWatched, a.Anime.Title)
	}
	// Output:
	// ID:   967, Status:      "watching", Episodes Watched:  73 Hokuto no Ken
	// ID:   820, Status:      "watching", Episodes Watched:   1 Ginga Eiyuu Densetsu
	// ID: 42897, Status:      "watching", Episodes Watched:   2 Horimiya
	// ID:  1453, Status:      "watching", Episodes Watched:  28 Maison Ikkoku
	// ID: 37521, Status:     "completed", Episodes Watched:  24 Vinland Saga
}

func ExampleSite_User_updatemyliststatus() {
	ctx := context.Background()

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()

	c, err := mal.NewSite(nil, &server.URL)
	if err != nil {
		fmt.Printf("Site creation error: %v", err)
		return
	}

	opts := c.Anime.UpdateMyListStatusOptions
	s, _, err := c.Anime.UpdateMyListStatus(ctx, 967,
		opts.AnimeStatus.Watching(),
		opts.NumEpisodesWatched(73),
		opts.Score(8),
		opts.Comments("You wa shock!"),
		opts.StartDate(time.Date(2022, 02, 20, 0, 0, 0, 0, time.UTC)),
		opts.FinishDate(time.Time{}), // Remove an existing date.
	)
	if err != nil {
		fmt.Printf("Anime.UpdateMyListStatus error: %v", err)
		return
	}
	fmt.Printf("Status: %q, Score: %d, Episodes Watched: %d, Comments: %q, Start Date: %s\n", s.Status, s.Score, s.NumEpisodesWatched, s.Comments, s.StartDate)
	// Output:
	// Status: "watching", Score: 8, Episodes Watched: 73, Comments: "You wa shock!", Start Date: 2022-02-20
}
