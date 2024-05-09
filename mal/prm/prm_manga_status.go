package prm

import "net/url"

// MangaStatus is an option that allows to filter the returned manga list by the
// specified status when using the UserService.MangaList method. It can also be
// passed as an option when updating the manga list.
type MangaStatus string

const (
	// MangaStatusReading returns the manga with status 'reading' from a user's
	// list or sets the status of a list item as such.
	MangaStatusReading MangaStatus = "reading"
	// MangaStatusCompleted returns the manga with status 'completed' from a
	// user's list or sets the status of a list item as such.
	MangaStatusCompleted MangaStatus = "completed"
	// MangaStatusOnHold returns the manga with status 'on hold' from a user's
	// list or sets the status of a list item as such.
	MangaStatusOnHold MangaStatus = "on_hold"
	// MangaStatusDropped returns the manga with status 'dropped' from a user's
	// list or sets the status of a list item as such.
	MangaStatusDropped MangaStatus = "dropped"
	// MangaStatusPlanToRead returns the manga with status 'plan to read' from a
	// user's list or sets the status of a list item as such.
	MangaStatusPlanToRead MangaStatus = "plan_to_read"
)

func (s MangaStatus) MangaListApply(v *url.Values)               { v.Set("status", string(s)) }
func (s MangaStatus) UpdateMyMangaListStatusApply(v *url.Values) { v.Set("status", string(s)) }
