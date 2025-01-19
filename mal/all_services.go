package mal

import (
	"net/http"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_client"
)

// Site manages communication with the MyAnimeList API.
type Site struct {
	client *mal_client.Client

	Anime *AnimeService
	Manga *MangaService
	User  *UserService
	Forum *ForumService
}

// NewSite returns a new MyAnimeList API client. The httpClient parameter
// allows to specify the http.client that will be used for all API requests. If
// a nil httpClient is provided, a new http.Site will be used.
//
// In the typical case, you will want to provide an http.Site that will
// perform the authentication for you. Such a client is provided by the
// golang.org/x/oauth2 package. Check out the example directory of the project
// for a full authentication example.
func NewSite(httpClient *http.Client, baseURL *string) (*Site, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	if baseURL == nil {
		defaultURL := mal_client.DefaultBaseURL
		baseURL = &defaultURL
	}

	baseRelURL, err := url.Parse(*baseURL)
	if err != nil {
		return nil, err
	}

	c := mal_client.NewClient(httpClient, baseRelURL)
	return &Site{
		client: c,
		User:   NewUserService(c),
		Anime:  NewAnimeService(c),
		Manga:  NewMangaService(c),
		Forum:  NewForumService(c),
	}, nil
}
