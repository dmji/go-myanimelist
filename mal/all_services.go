package mal

import (
	"context"
	"net/http"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/api_driver"
)

// Site manages communication with the MyAnimeList API.
type Site struct {
	client *api_driver.Client

	Anime *AnimeService
	Manga *MangaService
	User  *UserService
	Forum *ForumService
}

// NewClient returns a new MyAnimeList API client. The httpClient parameter
// allows to specify the http.client that will be used for all API requests. If
// a nil httpClient is provided, a new http.Site will be used.
//
// In the typical case, you will want to provide an http.Site that will
// perform the authentication for you. Such a client is provided by the
// golang.org/x/oauth2 package. Check out the example directory of the project
// for a full authentication example.
func NewSite(httpClient *http.Client) *Site {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, err := url.Parse(api_driver.DefaultBaseURL)
	if err != nil {
		panic(err)
	}

	c := &Site{
		client: api_driver.NewClient(httpClient, baseURL),
	}

	c.User = NewUserService(c.client)
	c.Anime = NewAnimeService(c.client)
	c.Manga = NewMangaService(c.client)
	c.Forum = NewForumService(c.client)

	return c
}

type HttpDriver interface {
	Do(ctx context.Context, req *http.Request, v interface{}) (*api_driver.Response, error)
	NewRequest(method, urlStr string, urlOptions ...func(v *url.Values)) (*http.Request, error)
}

func (c *Site) DirectRequest() HttpDriver {
	return c.client
}

func (c *Site) BaseURL() string {
	if c.client == nil {
		panic("client is nil")
	}
	return c.client.BaseURL.String()
}

func (c *Site) SetBaseURL(baseUrl *url.URL) {
	if c.client == nil {
		panic("client is nil")
	}

	c.client.BaseURL = baseUrl
}
