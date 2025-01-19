package mal_type

// The Forum of MyAnimeList.
type Forum struct {
	Categories []ForumCategory `json:"categories"`
}

// ForumCategory is a category of the forum.
type ForumCategory struct {
	Title  string       `json:"title"`
	Boards []ForumBoard `json:"boards"`
}

// ForumBoard is a board of the forum.
type ForumBoard struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Subboards   []ForumSubboard `json:"subboards"`
}

// ForumSubboard is a subboard of the forum.
type ForumSubboard struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
