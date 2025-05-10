package mal_prm

//go:generate go run github.com/dmji/go-stringer@latest -type=AnimeStatus,MangaStatus -trimprefix=@me -output status_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower

type AnimeStatus int8

const (
	// AnimeStatusWatching returns the anime with status 'watching' from a
	// user's list or sets the status of a list item as such.
	AnimeStatusWatching AnimeStatus = iota // "watching"
	// AnimeStatusCompleted returns the anime with status 'completed' from a
	// user's list or sets the status of a list item as such.
	AnimeStatusCompleted // "completed"
	// AnimeStatusOnHold returns the anime with status 'on hold' from a user's
	// list or sets the status of a list item as such.
	AnimeStatusOnHold // "on_hold"
	// AnimeStatusDropped returns the anime with status 'dropped' from a user's
	// list or sets the status of a list item as such.
	AnimeStatusDropped // "dropped"
	// AnimeStatusPlanToWatch returns the anime with status 'plan to watch' from
	// a user's list or sets the status of a list item as such.
	AnimeStatusPlanToWatch // "plan_to_watch"
)

// MangaStatus is an option that allows to filter the returned manga list by the
// specified status when using the UserService.MangaList method. It can also be
// passed as an option when updating the manga list.
type MangaStatus int8

const (
	// MangaStatusReading returns the manga with status 'reading' from a user's
	// list or sets the status of a list item as such.
	MangaStatusReading MangaStatus = iota // "reading"
	// MangaStatusCompleted returns the manga with status 'completed' from a
	// user's list or sets the status of a list item as such.
	MangaStatusCompleted // "completed"
	// MangaStatusOnHold returns the manga with status 'on hold' from a user's
	// list or sets the status of a list item as such.
	MangaStatusOnHold // "on_hold"
	// MangaStatusDropped returns the manga with status 'dropped' from a user's
	// list or sets the status of a list item as such.
	MangaStatusDropped // "dropped"
	// MangaStatusPlanToRead returns the manga with status 'plan to read' from a
	// user's list or sets the status of a list item as such.
	MangaStatusPlanToRead // "plan_to_read"
)
