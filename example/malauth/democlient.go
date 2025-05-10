package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal_prm"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// demoClient has methods showcasing the usage of the different MyAnimeList API
// methods. It stores the first error it encounters so error checking only needs
// to be done once.
//
// This pattern is used for convenience and should not be used in concurrent
// code without guarding the error.
type demoClient struct {
	*mal.Site
	err error
}

func (c *demoClient) showcase(ctx context.Context) error {
	methods := []func(context.Context){
		// Uncomment the methods you need to see their results. Run or build
		// using -tags=debug to see the full HTTP request and response.
		c.userMyInfo,
		/* 		c.animeList,
		   		c.mangaList,
		   		c.animeDetails,
		   		c.mangaDetails,
		   		c.animeRanking,
		   		c.mangaRanking,
		   		c.animeSeasonal,
		   		c.animeSuggested,
		   		c.animeListForLoop, // Warning: Many requests.
		   		c.updateMyAnimeListStatus,
		   		c.userAnimeList,
		   		c.deleteMyAnimeListItem,
		   		c.updateMyMangaListStatus,
		   		c.userMangaList,
		   		c.deleteMyMangaListItem,
		   		c.forumBoards,
		   		c.forumTopics,
		   		c.forumTopicDetails, */
	}
	for _, m := range methods {
		m(ctx)
	}
	if c.err != nil {
		return c.err
	}
	return nil
}

func (c *demoClient) userMyInfo(ctx context.Context) {
	if c.err != nil {
		return
	}
	u, _, err := c.User.MyInfo(ctx, nil)
	if err != nil {
		c.err = err
		return
	}

	fmt.Printf("ID: %5d, Joined: %v, Username: %s\n", u.ID, u.JoinedAt.Format("Jan 2006"), u.Name)
}

func (c *demoClient) animeList(ctx context.Context) {
	if c.err != nil {
		return
	}

	opts := c.Anime.ListOptions
	opts.AnimeFields.MyListStatus()
	anime, _, err := c.Anime.List(ctx, "hokuto no ken",
		opts.Fields(
			opts.AnimeFields.Rank(),
			opts.AnimeFields.Popularity(),
			opts.AnimeFields.StartSeason(),
		),
		opts.Limit(3),
		opts.Offset(0),
	)
	if err != nil {
		c.err = err
		return
	}
	for _, a := range anime {
		fmt.Printf("ID: %5d, Rank: %5d, Popularity: %5d %s (%d)\n", a.ID, a.Rank, a.Popularity, a.Title, a.StartSeason.Year)
	}
}

func (c *demoClient) mangaList(ctx context.Context) {
	if c.err != nil {
		return
	}

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
		c.err = err
		return
	}
	for _, m := range manga {
		fmt.Printf("ID: %5d, Volumes: %3d, Chapters: %3d %s (%s)\n", m.ID, m.NumVolumes, m.NumChapters, m.Title, m.AlternativeTitles.En)
	}
}

func (c *demoClient) animeDetails(ctx context.Context) {
	if c.err != nil {
		return
	}

	a, _, err := c.Anime.Details(ctx, 967,
		&mal_prm.AnimeDetailsRequestParameters{
			Fields: []mal_prm.AnimeField{
				mal_prm.AnimeFieldTypeAlternativeTitles.AnimeField(),
				mal_prm.AnimeFieldTypeMediaType.AnimeField(),
				mal_prm.AnimeFieldTypeNumEpisodes.AnimeField(),
				mal_prm.AnimeFieldTypeStartSeason.AnimeField(),
				mal_prm.AnimeFieldTypeSource.AnimeField(),
				mal_prm.AnimeFieldTypeGenres.AnimeField(),
				mal_prm.AnimeFieldTypeStudios.AnimeField(),
				mal_prm.AnimeFieldTypeAverageEpisodeDuration.AnimeField(),
			},
		},
	)
	if err != nil {
		c.err = err
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
}

func (c *demoClient) mangaDetails(ctx context.Context) {
	if c.err != nil {
		return
	}

	m, _, err := c.Manga.Details(ctx, 401,
		&mal_prm.MangaDetailsRequestParameters{
			Fields: []mal_prm.MangaField{
				mal_prm.MangaFieldTypeAlternativeTitles.MangaField(),
				mal_prm.MangaFieldTypeMediaType.MangaField(),
				mal_prm.MangaFieldTypeNumVolumes.MangaField(),
				mal_prm.MangaFieldTypeNumChapters.MangaField(),
				mal_prm.MangaFieldTypeAuthors.MangaField("last_name", "first_name"),
				mal_prm.MangaFieldTypeGenres.MangaField(),
				mal_prm.MangaFieldTypeStatus.MangaField(),
			},
		},
	)
	if err != nil {
		c.err = err
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
}

func (c *demoClient) animeListForLoop(ctx context.Context) {
	if c.err != nil {
		return
	}
	offset := 0
	opts := c.Anime.ListOptions
	for {
		anime, resp, err := c.Anime.List(ctx, "kiseijuu",
			opts.Fields(
				opts.AnimeFields.Rank(),
				opts.AnimeFields.Popularity(),
				opts.AnimeFields.StartSeason(),
			),
			opts.Limit(100),
			opts.Offset(offset),
		)
		if err != nil {
			c.err = err
			return
		}
		for _, a := range anime {
			fmt.Printf("ID: %5d, Rank: %5d, Popularity: %5d %s (%d)\n", a.ID, a.Rank, a.Popularity, a.Title, a.StartSeason.Year)
		}
		fmt.Println("--------")
		fmt.Printf("Next offset: %d\n", resp.NextOffset)
		offset = resp.NextOffset
		if offset == 0 {
			break
		}
	}
}

func (c *demoClient) userAnimeList(ctx context.Context) {
	if c.err != nil {
		return
	}

	anime, _, err := c.User.AnimeList(ctx, "@me",
		&mal_prm.UserAnimeListRequestParameters{
			Fields: []mal_prm.AnimeField{
				mal_prm.AnimeFieldTypeListStatus.AnimeField(),
			},
			Limit:  5,
			Sort:   mal_prm.SortAnimeListByListUpdatedAt,
			Status: mal_prm.AnimeStatusWatching,
		},
	)
	if err != nil {
		c.err = err
		return
	}
	for _, a := range anime {
		fmt.Printf("ID: %5d, Status: %15q, Episodes Watched: %3d %s\n", a.Anime.ID, a.Status.Status, a.Status.NumEpisodesWatched, a.Anime.Title)
	}
}

func (c *demoClient) userMangaList(ctx context.Context) {
	if c.err != nil {
		return
	}

	manga, _, err := c.User.MangaList(ctx, "@me",
		&mal_prm.UserMangaListRequestParameters{
			Fields: []mal_prm.MangaField{
				mal_prm.MangaFieldTypeListStatus.MangaField("comments", "tags"),
			},
			Limit:  5,
			Offset: 1,
			Sort:   mal_prm.SortMangaListByListScore,
		},
	)
	if err != nil {
		c.err = err
		return
	}
	for _, m := range manga {
		fmt.Printf("ID: %5d, Status: %15q, Volumes Read: %3d, Chapters Read: %3d %s\n", m.Manga.ID, m.Status.Status, m.Status.NumVolumesRead, m.Status.NumChaptersRead, m.Manga.Title)
	}
}

func (c *demoClient) updateMyAnimeListStatus(ctx context.Context) {
	if c.err != nil {
		return
	}

	opts := c.Anime.UpdateMyListStatusOptions
	s, _, err := c.Anime.UpdateMyListStatus(ctx, 967,
		opts.AnimeStatus.Watching(),
		opts.NumEpisodesWatched(73),
		opts.Score(8),
		opts.Comments("You wa shock!"),
		opts.StartDate(time.Date(2022, 0o2, 20, 0, 0, 0, 0, time.UTC)),
		opts.FinishDate(time.Time{}), // Remove an existing date.
	)
	if err != nil {
		c.err = err
		return
	}
	fmt.Printf("Status: %q, Score: %d, Episodes Watched: %d, Comments: %q, Start Date: %s\n", s.Status, s.Score, s.NumEpisodesWatched, s.Comments, s.StartDate)
}

func (c *demoClient) updateMyMangaListStatus(ctx context.Context) {
	if c.err != nil {
		return
	}

	opts := c.Manga.UpdateMyListStatusOptions
	s, _, err := c.Manga.UpdateMyListStatus(ctx, 401,
		opts.MangaStatus.Reading(),
		opts.NumVolumesRead(1),
		opts.NumChaptersRead(5),
		opts.Comments("Migi"),
		opts.StartDate(time.Date(2022, 0o2, 20, 0, 0, 0, 0, time.UTC)),
		opts.FinishDate(time.Time{}), // Remove an existing date.
	)
	if err != nil {
		c.err = err
		return
	}
	fmt.Printf("Status: %q, Volumes Read: %d, Chapters Read: %d, Comments: %q, Start Date: %s\n", s.Status, s.NumVolumesRead, s.NumChaptersRead, s.Comments, s.StartDate)
}

func (c *demoClient) deleteMyAnimeListItem(ctx context.Context) {
	if c.err != nil {
		return
	}
	_, err := c.Anime.DeleteMyListItem(ctx, 820)
	if err != nil {
		c.err = err
		return
	}
}

func (c *demoClient) deleteMyMangaListItem(ctx context.Context) {
	if c.err != nil {
		return
	}
	_, err := c.Manga.DeleteMyListItem(ctx, 1)
	if err != nil {
		c.err = err
		return
	}
}

func (c *demoClient) animeRanking(ctx context.Context) {
	if c.err != nil {
		return
	}

	rankings := []mal_prm.AnimeRanking{
		mal_prm.AnimeRankingAiring,
		mal_prm.AnimeRankingAll,
		mal_prm.AnimeRankingByPopularity,
	}

	opts := c.Anime.RankingOptions
	for _, r := range rankings {
		fmt.Println("Ranking:", r)
		anime, _, err := c.Anime.Ranking(ctx, r,
			opts.Fields(
				opts.AnimeFields.Rank(),
				opts.AnimeFields.Popularity(),
			),
		)
		if err != nil {
			c.err = err
			return
		}
		for _, a := range anime {
			fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
		}
		fmt.Println("--------")
	}
}

func (c *demoClient) mangaRanking(ctx context.Context) {
	if c.err != nil {
		return
	}

	opts := c.Manga.RankingOptions
	manga, _, err := c.Manga.Ranking(ctx,
		mal_prm.MangaRankingByPopularity,
		opts.Fields(
			opts.AnimeFields.Rank(),
			opts.AnimeFields.Popularity(),
		),
		opts.Limit(6),
	)
	if err != nil {
		c.err = err
		return
	}
	for _, m := range manga {
		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", m.Rank, m.Popularity, m.Title)
	}
}

func (c *demoClient) animeSeasonal(ctx context.Context) {
	if c.err != nil {
		return
	}

	opts := c.Anime.SeasonalOptions
	anime, _, err := c.Anime.Seasonal(ctx, 2020, opts.AnimeSeason.Fall(),
		opts.Fields(
			opts.AnimeFields.Rank(),
			opts.AnimeFields.Popularity(),
		),
		opts.SortSeasonalAnime.ByUsersCount(),
		opts.Limit(3),
		opts.Offset(0),
	)
	if err != nil {
		c.err = err
		return
	}
	for _, a := range anime {
		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
	}
}

func (c *demoClient) animeSuggested(ctx context.Context) {
	if c.err != nil {
		return
	}

	opts := c.Anime.SuggestedOptions
	anime, _, err := c.Anime.Suggested(ctx,
		opts.Limit(3),
		opts.Fields(
			opts.AnimeFields.Rank(),
			opts.AnimeFields.Popularity(),
		),
	)
	if err != nil {
		c.err = err
		return
	}
	for _, a := range anime {
		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
	}
}

func (c *demoClient) forumBoards(ctx context.Context) {
	if c.err != nil {
		return
	}
	forum, _, err := c.Forum.Boards(ctx)
	if err != nil {
		c.err = err
		return
	}
	for _, category := range forum.Categories {
		fmt.Printf("%s\n", category.Title)
		for _, b := range category.Boards {
			fmt.Printf("|-> %s\n", b.Title)
			for _, b := range b.Subboards {
				fmt.Printf("    |-> %s\n", b.Title)
			}
		}
		fmt.Println("---")
	}
}

func (c *demoClient) forumTopics(ctx context.Context) {
	if c.err != nil {
		return
	}

	topics, _, err := c.Forum.Topics(ctx,
		&mal_prm.ForumTopicsRequestParameters{
			Query: "JoJo opening",
			Limit: 2,
			Sort:  mal_prm.SortTopicsRecent,
		},
	)
	if err != nil {
		c.err = err
		return
	}
	for _, t := range topics {
		fmt.Printf("ID: %5d, Title: %5q created by %q\n", t.ID, t.Title, t.CreatedBy.Name)
	}
}

func (c *demoClient) forumTopicDetails(ctx context.Context) {
	if c.err != nil {
		return
	}
	topicDetails, _, err := c.Forum.TopicDetails(ctx,
		1877721, &mal_prm.ForumTopicDetailsRequestParameters{
			Limit: 3,
		},
	)
	if err != nil {
		c.err = err
		return
	}
	fmt.Printf("Topic title: %q\n", topicDetails.Title)
	if topicDetails.Poll != nil {
		fmt.Printf("Poll: %q\n", topicDetails.Poll.Question)
		for _, o := range topicDetails.Poll.Options {
			fmt.Printf("- %-25s %2d\n", o.Text, o.Votes)
		}
	}
	for _, p := range topicDetails.Posts {
		fmt.Printf("Post: %2d created by %q\n", p.Number, p.CreatedBy.Name)
	}
}
