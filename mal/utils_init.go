package mal

import (
	"context"
	"net/http"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_prm"
	"github.com/dmji/go-myanimelist/mal_type"
)

type client interface {
	// anime
	RequestGet(ctx context.Context, path string, v interface{}, q interface{}) (*mal_client.Response, error)
	UpdateMyListStatus(ctx context.Context, path string, id int, v interface{}, options ...func(v *url.Values)) (*mal_client.Response, error)
	DeleteMyListItem(ctx context.Context, path string, animeID int) (*mal_client.Response, error)
	RequestAnimeList(ctx context.Context, path string, opts *mal_prm.UserAnimeListRequestParameters) ([]mal_type.UserAnime, *mal_client.Response, error)
	// clientAnime

	// maga
	// RequestGet(ctx context.Context, path string, v interface{}, q interface{}) (*mal_client.Response, error)
	// UpdateMyListStatus(ctx context.Context, path string, id int, v interface{}, opts *mal_prm.UserMangaListRequestParameters) (*mal_client.Response, error)
	// DeleteMyListItem(ctx context.Context, path string, animeID int) (*mal_client.Response, error)
	RequestMangaList(ctx context.Context, path string, opts *mal_prm.UserMangaListRequestParameters) ([]mal_type.UserManga, *mal_client.Response, error)
	// clientManga

	// user
	// RequestGet(ctx context.Context, path string, v interface{}, q interface{}) (*mal_client.Response, error)
	// RequestMangaList(ctx context.Context, path string, opts *mal_prm.UserMangaListRequestParameters) ([]mal_type.UserManga, *mal_client.Response, error)
	// RequestAnimeList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserAnime, *mal_client.Response, error)
	// clientUser

	// forum
	// RequestGet(ctx context.Context, path string, v interface{}, q interface{}) (*mal_client.Response, error)
	RequestTopicDetails(ctx context.Context, path string, opts *mal_prm.ForumTopicDetailsRequestParameters) (mal_type.TopicDetails, *mal_client.Response, error)
	RequestTopics(ctx context.Context, path string, opts *mal_prm.ForumTopicsRequestParameters) ([]mal_type.Topic, *mal_client.Response, error)
	clientForum
}

type initOptions struct {
	c client
}

func (o *initOptions) initEmptyFields() error {
	var err error
	if o.c == nil {
		o.c, err = mal_client.NewClientUrl(nil, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

type fnOptionApply func(*initOptions) error

func WithCustomClientPtr(c client) fnOptionApply {
	return func(o *initOptions) error {
		o.c = c
		return nil
	}
}

func WithCustomClientUrl(httpClient *http.Client, baseURL *string) fnOptionApply {
	return func(o *initOptions) error {
		var err error
		o.c, err = mal_client.NewClientUrl(httpClient, baseURL)
		return err
	}
}

func WithCustomClient(httpClient *http.Client, baseURL *url.URL) fnOptionApply {
	return func(o *initOptions) error {
		o.c = mal_client.NewClient(httpClient, baseURL)
		return nil
	}
}
