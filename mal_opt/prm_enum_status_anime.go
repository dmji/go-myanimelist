package mal_opt

import "net/url"

// AnimeStatus is an option that allows to filter the returned anime list by the
// specified status when using the UserService.AnimeList method. It can also be
// passed as an option when updating the anime list.
type AnimeStatus string

const (
	// AnimeStatusWatching returns the anime with status 'watching' from a
	// user's list or sets the status of a list item as such.
	AnimeStatusWatching AnimeStatus = "watching"
	// AnimeStatusCompleted returns the anime with status 'completed' from a
	// user's list or sets the status of a list item as such.
	AnimeStatusCompleted AnimeStatus = "completed"
	// AnimeStatusOnHold returns the anime with status 'on hold' from a user's
	// list or sets the status of a list item as such.
	AnimeStatusOnHold AnimeStatus = "on_hold"
	// AnimeStatusDropped returns the anime with status 'dropped' from a user's
	// list or sets the status of a list item as such.
	AnimeStatusDropped AnimeStatus = "dropped"
	// AnimeStatusPlanToWatch returns the anime with status 'plan to watch' from
	// a user's list or sets the status of a list item as such.
	AnimeStatusPlanToWatch AnimeStatus = "plan_to_watch"
)

func (s AnimeStatus) UpdateMyAnimeListStatusApply(v *url.Values) { v.Set("status", string(s)) }

func (n AnimeStatus) Watching() AnimeStatus    { return AnimeStatusWatching }
func (n AnimeStatus) Completed() AnimeStatus   { return AnimeStatusCompleted }
func (n AnimeStatus) OnHold() AnimeStatus      { return AnimeStatusOnHold }
func (n AnimeStatus) Dropped() AnimeStatus     { return AnimeStatusDropped }
func (n AnimeStatus) PlanToWatch() AnimeStatus { return AnimeStatusPlanToWatch }
