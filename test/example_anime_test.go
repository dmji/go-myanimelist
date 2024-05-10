package mal_test

import (
	"context"
	_ "embed"
	"fmt"
	"net/url"
	"strings"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal/prm"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ExampleSite_Anime_list() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	anime, _, err := c.Anime.List(ctx, "hokuto no ken",
		prm.Fields{"rank", "popularity", "start_season"},
		prm.Limit(5),
		prm.Offset(0),
	)
	if err != nil {
		fmt.Printf("Anime.List error: %v", err)
		return
	}
	for _, a := range anime {
		fmt.Printf("ID: %5d, Rank: %5d, Popularity: %5d %s (%d)\n", a.ID, a.Rank, a.Popularity, a.Title, a.StartSeason.Year)
	}
	// Output:
	// ID:   967, Rank:   556, Popularity:  1410 Hokuto no Ken (1984)
	// ID:  1356, Rank:  1423, Popularity:  3367 Hokuto no Ken 2 (1987)
	// ID:  1358, Rank:  2757, Popularity:  3964 Hokuto no Ken Movie (1986)
}

func ExampleSite_Anime_details() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	a, _, err := c.Anime.Details(ctx, 967,
		prm.Fields{
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
	if err != nil {
		fmt.Printf("Anime.Details error: %v", err)
		return
	}

	title := cases.Title(language.Und)
	fmt.Printf("%s\n", a.Title)
	fmt.Printf("ID: %d\n", a.ID)
	fmt.Printf("English: %s\n", a.AlternativeTitles.En)
	fmt.Printf("Type: %s\n", strings.ToUpper(a.MediaType))
	fmt.Printf("Episodes: %d\n", a.NumEpisodes)
	fmt.Printf("Premiered: %s %d\n", title.String(a.StartSeason.Season), a.StartSeason.Year)
	fmt.Print("Studios: ")
	delim := ""
	for _, s := range a.Studios {
		fmt.Printf("%s%s", delim, s.Name)
		delim = " "
	}
	fmt.Println()
	fmt.Printf("Source: %s\n", title.String(a.Source))
	fmt.Print("Genres: ")
	delim = ""
	for _, g := range a.Genres {
		fmt.Printf("%s%s", delim, g.Name)
		delim = " "
	}
	fmt.Println()
	fmt.Printf("Duration: %d min. per ep.\n", a.AverageEpisodeDuration/60)
	// Output:
	// Hokuto no Ken
	// ID: 967
	// English: Fist of the North Star
	// Type: TV
	// Episodes: 109
	// Premiered: Fall 1984
	// Studios: Toei Animation
	// Source: Manga
	// Genres: Action Drama Martial Arts Sci-Fi Shounen
	// Duration: 25 min. per ep.
}

func ExampleSite_Anime_ranking() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	anime, _, err := c.Anime.Ranking(ctx,
		prm.AnimeRankingAiring,
		prm.Fields{"rank", "popularity"},
		prm.Limit(6),
	)
	if err != nil {
		fmt.Printf("Anime.Ranking error: %v", err)
		return
	}
	for _, a := range anime {
		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
	}
	// Output:
	// Rank:     2, Popularity:   104 Shingeki no Kyojin: The Final Season
	// Rank:    59, Popularity:   371 Re:Zero kara Hajimeru Isekai Seikatsu 2nd Season Part 2
	// Rank:    67, Popularity:  1329 Yuru Camp△ Season 2
	// Rank:    69, Popularity:   109 Jujutsu Kaisen (TV)
	// Rank:    83, Popularity:  3714 Holo no Graffiti
	// Rank:    85, Popularity:   304 Horimiya
}

func ExampleSite_Anime_seasonal() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	opts := prm.SeasonalAnimeOptionProvider{}
	anime, _, err := c.Anime.Seasonal(ctx, 2020, opts.AnimeSeason.Fall(),
		prm.Fields{"rank", "popularity"},
		opts.SortSeasonalAnime.ByUsersCount(),
		opts.Limit.Val(3),
		opts.Offset.Val(0),
	)
	if err != nil {
		fmt.Printf("Anime.Seasonal error: %v", err)
		return
	}
	for _, a := range anime {
		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
	}
	// Output:
	// Rank:    93, Popularity:    31 One Piece
	// Rank:  1870, Popularity:    92 Black Clover
	// Rank:    62, Popularity:   106 Jujutsu Kaisen (TV)
}

func ExampleSite_Anime_suggested() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	anime, _, err := c.Anime.Suggested(ctx,
		prm.Limit(10),
		prm.Fields{"rank", "popularity"},
	)
	if err != nil {
		fmt.Printf("Anime.Suggested error: %v", err)
		return
	}
	for _, a := range anime {
		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
	}
	// Output:
	// Rank:   971, Popularity:   916 Kill la Kill Specials
	// Rank:   726, Popularity:  4972 Osomatsu-san Movie
	// Rank:   943, Popularity:  4595 Watashi no Ashinaga Ojisan
}

func ExampleSite_Anime_deletemylistttem() {
	ctx := context.Background()

	c := mal.NewSite(nil)

	// Ignore the 3 following lines. A stub server is used instead of the real
	// API to produce testable examples. See: https://go.dev/blog/examples
	server := newStubServer()
	defer server.Close()
	baseURL, _ := url.Parse(server.URL)
	c.SetBaseURL(baseURL)

	resp, err := c.Anime.DeleteMyListItem(ctx, 967)
	if err != nil {
		fmt.Printf("Anime.DeleteMyListItem error: %v", err)
		return
	}
	fmt.Println(resp.Status)
	// Output:
	// 200 OK
}
