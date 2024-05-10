package mal_test

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal/prm"
)

func ExampleSite_User_mangalist() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	manga, _, err := c.User.MangaList(ctx, "@me",
		prm.Fields{"list_status"},
		prm.SortMangaListByListUpdatedAt,
		prm.Limit(2),
	)
	if err != nil {
		fmt.Printf("User.MangaList error: %v", err)
		return
	}
	for _, m := range manga {
		fmt.Printf("ID: %5d, Status: %15q, Volumes Read: %3d, Chapters Read: %3d %s\n", m.Manga.ID, m.Status.Status, m.Status.NumVolumesRead, m.Status.NumChaptersRead, m.Manga.Title)
	}
	// Output:
	// ID:    21, Status:     "completed", Volumes Read:  12, Chapters Read: 108 Death Note
	// ID:   401, Status:       "reading", Volumes Read:   1, Chapters Read:   5 Kiseijuu
}

func ExampleSite_Manga_updatemyliststatus() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	s, _, err := c.Manga.UpdateMyListStatus(ctx, 401,
		prm.MangaStatusReading,
		prm.NumVolumesRead(1),
		prm.NumChaptersRead(5),
		prm.Comments("Migi"),
		prm.StartDate(time.Date(2022, 02, 20, 0, 0, 0, 0, time.UTC)),
		prm.FinishDate(time.Time{}), // Remove an existing date.
	)
	if err != nil {
		fmt.Printf("Manga.UpdateMyListStatus error: %v", err)
		return
	}
	fmt.Printf("Status: %q, Volumes Read: %d, Chapters Read: %d, Comments: %q, Start Date: %s\n", s.Status, s.NumVolumesRead, s.NumChaptersRead, s.Comments, s.StartDate)
	// Output:
	// Status: "reading", Volumes Read: 1, Chapters Read: 5, Comments: "Migi", Start Date: 2022-02-20
}
