package mal

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/dmji/go-myanimelist/mal/common"
)

// MARK: AnimeStatus
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

func (s AnimeStatus) AnimeListApply(v *url.Values)               { v.Set("status", string(s)) }
func (s AnimeStatus) updateMyAnimeListStatusApply(v *url.Values) { v.Set("status", string(s)) }

var itoa = strconv.Itoa

// MARK: Score
// Score is an option that can update the anime and manga list scores with
// values 0-10.
type Score int

func (s Score) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("score", itoa(int(s)))
}
func (s Score) updateMyMangaListStatusApply(v *url.Values) {
	v.Set("score", itoa(int(s)))
}

// MARK: NumEpisodesWatched
// NumEpisodesWatched is an option that can update the number of episodes
// watched of an anime in the user's list.
type NumEpisodesWatched int

func (n NumEpisodesWatched) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("num_watched_episodes", itoa(int(n)))
}

// MARK: NumTimesRewatched
// NumTimesRewatched is an option that can update the number of times the user
// has rewatched an anime in their list.
type NumTimesRewatched int

func (n NumTimesRewatched) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("num_times_rewatched", itoa(int(n)))
}

// MARK: IsRewatching
// IsRewatching is an option that can update if a user is rewatching an anime in
// their list.
type IsRewatching bool

func (r IsRewatching) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("is_rewatching", strconv.FormatBool(bool(r)))
}

// MARK: RewatchValue
// RewatchValue is an option that can update the rewatch value of an anime in
type RewatchValue int

const (
	RewatchNoValue  RewatchValue = 0
	RewatchVeryLow  RewatchValue = 1
	RewatchLow      RewatchValue = 2
	RewatchMedium   RewatchValue = 3
	RewatchHigh     RewatchValue = 4
	RewatchVeryHigh RewatchValue = 5
)

func (r RewatchValue) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("rewatch_value", itoa(int(r)))
}

// MARK: Priority
// Priority is an option that allows to update the priority of an anime or manga
// in the user's list with values 0=Low, 1=Medium, 2=High.
type Priority int

func (p Priority) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("priority", itoa(int(p)))
}
func (p Priority) updateMyMangaListStatusApply(v *url.Values) {
	v.Set("priority", itoa(int(p)))
}

// MARK: Tags
// Tags is an option that allows to update the comma-separated tags of anime and
// manga in the user's list.
type Tags []string

func (t Tags) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("tags", strings.Join(t, ","))
}
func (t Tags) updateMyMangaListStatusApply(v *url.Values) {
	v.Set("tags", strings.Join(t, ","))
}

// MARK: Comments
// Comments is an option that allows to update the comment of anime and manga in
// the user's list.
type Comments string

func (c Comments) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("comments", string(c))
}
func (c Comments) updateMyMangaListStatusApply(v *url.Values) {
	v.Set("comments", string(c))
}

// MARK: StartDate
// StartDate is an option that allows to update the start date of anime and manga
// in the user's list.
type StartDate time.Time

func (d StartDate) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("start_date", common.FormatMALDate(time.Time(d)))
}
func (d StartDate) updateMyMangaListStatusApply(v *url.Values) {
	v.Set("start_date", common.FormatMALDate(time.Time(d)))
}

// MARK: FinishDate
// FinishDate is an option that allows to update the finish date of anime and manga
// in the user's list.
type FinishDate time.Time

func (d FinishDate) updateMyAnimeListStatusApply(v *url.Values) {
	v.Set("finish_date", common.FormatMALDate(time.Time(d)))
}
func (d FinishDate) updateMyMangaListStatusApply(v *url.Values) {
	v.Set("finish_date", common.FormatMALDate(time.Time(d)))
}
