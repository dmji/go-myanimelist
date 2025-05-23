package mal_test

import (
	"context"
	"encoding/json"
	"flag"
	"reflect"
	"testing"
	"time"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/go-myanimelist/mal_opt"
	"github.com/dmji/go-myanimelist/mal_type"
	"golang.org/x/oauth2"
)

var (
	oauth2Token  = flag.String("oauth2-token", "", "MyAnimeList.net oauth2 token to use for integration tests in `JSON` format")
	clientID     = flag.String("client-id", "", "your registered MyAnimeList.net application client ID")
	clientSecret = flag.String("client-secret", "", "your registered MyAnimeList.net application client secret; optional if you chose App Type 'other'")
)

func setupIntegration(ctx context.Context, t *testing.T) (*mal.Site, error) {
	const tokenFormat = `
	{
		"token_type": "Bearer",
		"access_token": "yourAccessToken",
		"refresh_token": "yourRefreshToken",
		"expiry": "2021-06-01T16:12:56.1319122Z"
		}`

	if *oauth2Token == "" || *clientID == "" {
		t.Log("No oauth2 token or client ID provided.")
		t.Log("The integration tests are meant to be run with a dedicated test account with empty lists.")
		t.Log("To run the integration tests use: go test --client-id='<your client ID>' --oauth2-token='<your oauth2 token>'")
		t.Logf("The oauth2 token is expected to be in JSON format, example: %s", tokenFormat)
		t.Log(`Note: On some terminals you may need to escape the double quotes: --oauth2-token='{\"token_type\":\"Bearer\",...'`)
		t.Skip("Skipping integration tests.")
	}

	token := new(oauth2.Token)
	err := json.Unmarshal([]byte(*oauth2Token), token)
	if err != nil {
		t.Logf("The oauth2 token is expected to be in JSON format, example: %s", tokenFormat)
		t.Log(`Note: On some terminals you may need to escape the double quotes: --oauth2-token='{\"token_type\":\"Bearer\",...'`)
		t.Logf("failed to unmarshal oauth2 token: %v", err)
		t.Fatalf("input was:\n%q", *oauth2Token)
	}

	conf := &oauth2.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://myanimelist.net/v1/oauth2/authorize",
			TokenURL:  "https://myanimelist.net/v1/oauth2/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	return mal.NewSite(mal.WithCustomClientUrl(conf.Client(ctx, token), nil))
}

func TestIntegration(t *testing.T) {
	ctx := context.Background()
	client, err := setupIntegration(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	username := testGetUserInfo(ctx, t, client)
	t.Run("UpdateUserAnimeList", func(t *testing.T) {
		testUpdateUserAnimeList(ctx, t, client, username)
	})
	t.Run("UpdateUserMangaList", func(t *testing.T) {
		testUpdateUserMangaList(ctx, t, client, username)
	})
	t.Run("AnimeMethods", func(t *testing.T) {
		testAnimeMethods(ctx, t, client)
	})
	t.Run("MangaMethods", func(t *testing.T) {
		testMangaMethods(ctx, t, client)
	})
	t.Run("ForumMethods", func(t *testing.T) {
		testForumMethods(ctx, t, client)
	})
}

func testGetUserInfo(ctx context.Context, t *testing.T, client *mal.Site) (username string) {
	t.Helper()
	// Get user info to find the username.
	info, _, err := client.User.MyInfo(ctx, nil)
	if err != nil {
		t.Fatal("User.MyInfo returned err:", err)
	}

	username = info.Name
	t.Logf("Running integration tests on account with username: %q", username)
	return username
}

func testUpdateUserAnimeList(ctx context.Context, t *testing.T, client *mal.Site, username string) {
	// Get anime list for test account.
	const me = "@me"
	list, _, err := client.User.AnimeList(ctx, me)
	if err != nil {
		t.Fatalf("User.AnimeList(%q) returned err: %s", me, err)
	}

	// Being strict here. Anime list of test account must be empty.
	if len(list) != 0 {
		t.Fatalf("MyAnimeList.net test account %q is supposed to have 0 anime but has %d", username, len(list))
	}

	testAnimeIDs := []int{1, 5, 6, 7}
	// Clean up all anime at the end.
	defer func() {
		for _, id := range testAnimeIDs {
			if _, delErr := client.Anime.DeleteMyListItem(ctx, id); delErr != nil {
				t.Errorf("Anime.DeleteMyListItem(%d) returned err: %v", id, delErr)
			}
		}
	}()

	// Test adding some anime.
	for _, id := range testAnimeIDs {
		if _, _, err := client.Anime.UpdateMyListStatus(ctx, id,
			mal_opt.AnimeStatusWatching,
			mal_opt.Comments("test comment"),
			mal_opt.IsRewatching(true),
			mal_opt.NumEpisodesWatched(1),
			mal_opt.NumTimesRewatched(1),
			mal_opt.Priority(1),
			mal_opt.RewatchValue(1),
			mal_opt.Score(1),
			mal_opt.Tags{"foo", "bar"},
			mal_opt.StartDate(time.Date(2022, 0o2, 20, 0, 0, 0, 0, time.UTC)),
			mal_opt.FinishDate(time.Time{}),
		); err != nil {
			t.Fatalf("Anime.UpdateMyListStatus(%d) returned err: %v", id, err)
		}
	}

	// Get anime list of test account for a second time.
	list, _, err = client.User.AnimeList(ctx, me,
		mal_opt.Fields{"list_status{num_times_rewatched, rewatch_value, priority, comments, tags}"},
	)
	if err != nil {
		t.Fatalf("User.AnimeList(%q) after additions returned err: %s", me, err)
	}

	// And make sure it has the number of anime it's supposed to have.
	if got, want := len(list), len(testAnimeIDs); got != want {
		t.Fatalf("Test account Anime number after additions = %d, want %d", got, want)
	}

	// And that they all have been updated appropriately.
	for _, a := range list {
		want := mal_type.AnimeListStatus{
			Status:             mal_opt.AnimeStatusWatching,
			Score:              1,
			NumEpisodesWatched: 1,
			IsRewatching:       true,
			Priority:           1,
			NumTimesRewatched:  1,
			RewatchValue:       1,
			Tags:               []string{"foo", "bar"},
			Comments:           "test comment",
			StartDate:          "2022-02-20",
			FinishDate:         "",
		}
		a.Status.UpdatedAt = time.Time{}
		if got := a.Status; !reflect.DeepEqual(got, want) {
			t.Errorf("Anime ID: %d AnimeListStatus\nhave: %+v\nwant: %+v", a.Anime.ID, got, want)
		}
	}
}

func testUpdateUserMangaList(ctx context.Context, t *testing.T, client *mal.Site, username string) {
	// Get manga list for test account.
	const me = "@me"
	list, _, err := client.User.MangaList(ctx, me)
	if err != nil {
		t.Fatalf("User.MangaList(%q) returned err: %s", me, err)
	}

	// Being strict here. Manga list of test account must be empty.
	if len(list) != 0 {
		t.Fatalf("MyMangaList.net test account %q is supposed to have 0 manga but has %d", username, len(list))
	}

	testMangaIDs := []int{1, 2, 3, 4}
	// Clean up all manga at the end.
	defer func() {
		for _, id := range testMangaIDs {
			if _, delErr := client.Manga.DeleteMyListItem(ctx, id); delErr != nil {
				t.Errorf("Manga.DeleteMyListItem(%d) returned err: %v", id, delErr)
			}
		}
	}()

	// Test adding some manga.
	for _, id := range testMangaIDs {
		if _, _, err := client.Manga.UpdateMyListStatus(ctx, id,
			mal_opt.MangaStatusReading,
			mal_opt.Comments("test comment"),
			mal_opt.IsRereading(true),
			mal_opt.NumChaptersRead(1),
			mal_opt.NumVolumesRead(1),
			mal_opt.NumTimesReread(1),
			mal_opt.Priority(1),
			mal_opt.RereadValue(1),
			mal_opt.Score(1),
			mal_opt.Tags{"foo", "bar"},
			mal_opt.StartDate(time.Date(2022, 0o2, 20, 0, 0, 0, 0, time.UTC)),
			mal_opt.FinishDate(time.Time{}),
		); err != nil {
			t.Fatalf("Manga.UpdateMyListStatus(%d) returned err: %v", id, err)
		}
	}

	// Get manga list of test account for a second time.
	list, _, err = client.User.MangaList(ctx, me,
		mal_opt.Fields{"list_status{num_times_reread, reread_value, priority, comments, tags}"},
	)
	if err != nil {
		t.Fatalf("User.MangaList(%q) after additions returned err: %s", me, err)
	}

	// And make sure it has the number of manga it's supposed to have.
	if got, want := len(list), len(testMangaIDs); got != want {
		t.Fatalf("Test account Manga number after additions = %d, want %d", got, want)
	}

	// And that they all have been updated appropriately.
	for _, a := range list {
		want := mal_type.MangaListStatus{
			Status:          mal_opt.MangaStatusReading,
			Score:           1,
			NumChaptersRead: 1,
			NumVolumesRead:  1,
			IsRereading:     true,
			Priority:        1,
			NumTimesReread:  1,
			RereadValue:     1,
			Tags:            []string{"foo", "bar"},
			Comments:        "test comment",
			StartDate:       "2022-02-20",
			FinishDate:      "",
		}
		a.Status.UpdatedAt = time.Time{}
		if got := a.Status; !reflect.DeepEqual(got, want) {
			t.Errorf("Manga ID: %d MangaListStatus\nhave: %+v\nwant: %+v", a.Manga.ID, got, want)
		}
	}
}

func testAnimeMethods(ctx context.Context, t *testing.T, client *mal.Site) {
	list, _, err := client.Anime.List(ctx, "kiseijuu", mal_opt.Limit(2))
	if err != nil {
		t.Errorf("Anime.List returned error: %v", err)
	}
	if len(list) == 0 {
		t.Fatal("Anime.List returned 0 anime")
	}

	_, _, err = client.Anime.Details(ctx, list[0].ID)
	if err != nil {
		t.Errorf("Anime.Details returned error: %v", err)
	}

	_, _, err = client.Anime.Ranking(ctx, mal_opt.AnimeRankingAll, mal_opt.Limit(2))
	if err != nil {
		t.Errorf("Anime.Ranking returned error: %v", err)
	}

	opts := mal_opt.SeasonalAnimeOptionProvider{}
	_, _, err = client.Anime.Seasonal(ctx, 2020, opts.AnimeSeason.Winter(), opts.SortSeasonalAnime.ByUsersCount(), opts.Limit(2))
	if err != nil {
		t.Errorf("Anime.Seasonal returned error: %v", err)
	}

	_, _, err = client.Anime.Suggested(ctx, mal_opt.Fields{"rank", "popularity"}, mal_opt.Limit(2))
	if err != nil {
		t.Errorf("Anime.Suggested returned error: %v", err)
	}
}

func testMangaMethods(ctx context.Context, t *testing.T, client *mal.Site) {
	list, _, err := client.Manga.List(ctx, "kiseijuu", mal_opt.Limit(2))
	if err != nil {
		t.Errorf("Manga.List returned error: %v", err)
	}
	if len(list) == 0 {
		t.Fatal("Manga.List returned 0 anime")
	}

	_, _, err = client.Manga.Details(ctx, list[0].ID)
	if err != nil {
		t.Errorf("Manga.Details returned error: %v", err)
	}

	_, _, err = client.Manga.Ranking(ctx, mal_opt.MangaRankingAll, mal_opt.Limit(2))
	if err != nil {
		t.Errorf("Manga.Ranking returned error: %v", err)
	}
}

func testForumMethods(ctx context.Context, t *testing.T, client *mal.Site) {
	_, _, err := client.Forum.Boards(ctx)
	if err != nil {
		t.Errorf("Forum.Boards returned error: %v", err)
	}

	topics, _, err := client.Forum.Topics(ctx, mal_opt.Query("kiseijuu"), mal_opt.Limit(2))
	if err != nil {
		t.Errorf("Forum.Topics returned error: %v", err)
	}
	if len(topics) == 0 {
		t.Fatal("Forum.Topics returned 0 topics")
	}

	_, _, err = client.Forum.TopicDetails(ctx, topics[0].ID, mal_opt.Limit(2))
	if err != nil {
		t.Errorf("Forum.TopicDetails returned error: %v", err)
	}
}
