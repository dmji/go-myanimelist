package mal

// Site manages communication with the MyAnimeList API.
type Site struct {
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
func NewSite(opts ...fnOptionApply) (*Site, error) {
	opt := &initOptions{}
	for _, fn := range opts {
		if err := fn(opt); err != nil {
			return nil, err
		}
	}
	if err := opt.initEmptyFields(); err != nil {
		return nil, err
	}

	return &Site{
		User:  NewUserService(opt.c),
		Anime: NewAnimeService(opt.c),
		Manga: NewMangaService(opt.c),
		Forum: NewForumService(opt.c),
	}, nil
}
