package mal_test

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ExampleSite_Manga_list() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	opts := c.Manga.ListOptions
	manga, _, err := c.Manga.List(ctx, "parasyte",
		opts.Fields(
			opts.MangaFields.NumVolumes(),
			opts.MangaFields.NumChapters(),
			opts.MangaFields.AlternativeTitles(),
		),
		opts.Limit(3),
		opts.Offset(0),
	)
	if err != nil {
		fmt.Printf("Manga.List error: %v", err)
		return
	}
	for _, m := range manga {
		fmt.Printf("ID: %5d, Volumes: %3d, Chapters: %3d %s (%s)\n", m.ID, m.NumVolumes, m.NumChapters, m.Title, m.AlternativeTitles.En)
	}
	// Output:
	// ID:   401, Volumes:  10, Chapters:  64 Kiseijuu (Parasyte)
	// ID: 78789, Volumes:   1, Chapters:  12 Neo Kiseijuu (Neo Parasyte m)
	// ID: 80797, Volumes:   1, Chapters:  15 Neo Kiseijuu f (Neo Parasyte f)
}

func ExampleSite_Manga_details() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

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
	if err != nil {
		fmt.Printf("Manga.Details error: %v", err)
		return
	}

	title := cases.Title(language.Und)
	fmt.Printf("%s\n", m.Title)
	fmt.Printf("ID: %d\n", m.ID)
	fmt.Printf("English: %s\n", m.AlternativeTitles.En)
	fmt.Printf("Type: %s\n", title.String(m.MediaType))
	fmt.Printf("Volumes: %d\n", m.NumVolumes)
	fmt.Printf("Chapters: %d\n", m.NumChapters)
	fmt.Print("Studios: ")
	delim := ""
	for _, s := range m.Authors {
		fmt.Printf("%s%s, %s (%s)", delim, s.Person.LastName, s.Person.FirstName, s.Role)
		delim = " "
	}
	fmt.Println()
	fmt.Print("Genres: ")
	delim = ""
	for _, g := range m.Genres {
		fmt.Printf("%s%s", delim, g.Name)
		delim = " "
	}
	fmt.Println()
	fmt.Printf("Status: %s\n", title.String(m.Status))
	// Output:
	// Kiseijuu
	// ID: 401
	// English: Parasyte
	// Type: Manga
	// Volumes: 10
	// Chapters: 64
	// Studios: Iwaaki, Hitoshi (Story & Art)
	// Genres: Action Psychological Sci-Fi Drama Horror Seinen
	// Status: Finished
}

func ExampleSite_Manga_ranking() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	opts := c.Manga.RankingOptions
	manga, _, err := c.Manga.Ranking(ctx,
		opts.MangaRanking.ByPopularity(),
		opts.Fields("rank", "popularity"),
		opts.Limit(6),
	)
	if err != nil {
		fmt.Printf("Manga.Ranking error: %v", err)
		return
	}
	for _, m := range manga {
		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", m.Rank, m.Popularity, m.Title)
	}
	// Output:
	// Rank:    38, Popularity:     1 Shingeki no Kyojin
	// Rank:     3, Popularity:     2 One Piece
	// Rank:     1, Popularity:     3 Berserk
	// Rank:   566, Popularity:     4 Naruto
	// Rank:   106, Popularity:     5 Tokyo Ghoul
	// Rank:    39, Popularity:     6 One Punch-Man
}

func ExampleSite_Manga_deletemylistitem() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	resp, err := c.Manga.DeleteMyListItem(ctx, 401)
	if err != nil {
		fmt.Printf("Manga.DeleteMyListItem error: %v", err)
		return
	}
	fmt.Println(resp.Status)
	// Output:
	// 200 OK
}
