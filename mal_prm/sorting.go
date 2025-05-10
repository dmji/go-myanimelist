package mal_prm

//go:generate go run github.com/dmji/go-stringer@latest -type=SortAnimeListBy,SortMangaListBy,SortTopics -trimprefix=@me -output sorting_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower

// SortAnimeListBy is an option that sorts the results when getting the user's
// anime list.
type SortAnimeListBy int8

const (
	// SortAnimeListByListScore sorts results by the score of each item in the
	// list in descending order.
	SortAnimeListByListScore SortAnimeListBy = iota // "list_score"
	// SortAnimeListByListUpdatedAt sorts results by the most updated entries in
	// the list in descending order.
	SortAnimeListByListUpdatedAt // "list_updated_at"
	// SortAnimeListByAnimeTitle sorts results by the anime title in ascending
	// order.
	SortAnimeListByAnimeTitle // "anime_title"
	// SortAnimeListByAnimeStartDate sorts results by the anime start date in
	// descending order.
	SortAnimeListByAnimeStartDate // "anime_start_date"
	// SortAnimeListByAnimeID sorts results by the anime ID in ascending order.
	// Note: Currently under development.
	SortAnimeListByAnimeID // "anime_id"
)

// SortMangaListBy is an option that sorts the results when getting the user's
// manga list.
type SortMangaListBy int8

const (
	// SortMangaListByListScore sorts results by the score of each item in the
	// list in descending order.
	SortMangaListByListScore SortMangaListBy = iota // "list_score"
	// SortMangaListByListUpdatedAt sorts results by the most updated entries in
	// the list in descending order.
	SortMangaListByListUpdatedAt // "list_updated_at"
	// SortMangaListByMangaTitle sorts results by the manga title in ascending
	// order.
	SortMangaListByMangaTitle // "manga_title"
	// SortMangaListByMangaStartDate sorts results by the manga start date in
	// descending order.
	SortMangaListByMangaStartDate // "manga_start_date"
	// SortMangaListByMangaID sorts results by the manga ID in ascending order.
	// Note: Currently under development.
	SortMangaListByMangaID // "manga_id"
)

type SortTopics int8

const (
	SortTopicsRecent SortTopics = iota
)
