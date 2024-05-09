package mal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/nstratos/go-myanimelist/mal/common"
)

// AnimeListOption are options specific to the UserService.AnimeList method.
type AnimeListOption interface {
	AnimeListApply(v *url.Values)
}

// SortAnimeList is an option that sorts the results when getting the user's
// anime list.
type SortAnimeList string

const (
	// SortAnimeListByListScore sorts results by the score of each item in the
	// list in descending order.
	SortAnimeListByListScore SortAnimeList = "list_score"
	// SortAnimeListByListUpdatedAt sorts results by the most updated entries in
	// the list in descending order.
	SortAnimeListByListUpdatedAt SortAnimeList = "list_updated_at"
	// SortAnimeListByAnimeTitle sorts results by the anime title in ascending
	// order.
	SortAnimeListByAnimeTitle SortAnimeList = "anime_title"
	// SortAnimeListByAnimeStartDate sorts results by the anime start date in
	// descending order.
	SortAnimeListByAnimeStartDate SortAnimeList = "anime_start_date"
	// SortAnimeListByAnimeID sorts results by the anime ID in ascending order.
	// Note: Currently under development.
	SortAnimeListByAnimeID SortAnimeList = "anime_id"
)

func (s SortAnimeList) AnimeListApply(v *url.Values) { v.Set("sort", string(s)) }

// UserAnime contains an anime record along with its status on the user's list.
type UserAnime struct {
	Anime  Anime           `json:"node"`
	Status AnimeListStatus `json:"list_status"`
}

// AnimeListStatus shows the status of each anime in a user's anime list.
type AnimeListStatus struct {
	Status             AnimeStatus `json:"status"`
	Score              int         `json:"score"`
	NumEpisodesWatched int         `json:"num_episodes_watched"`
	IsRewatching       bool        `json:"is_rewatching"`
	UpdatedAt          time.Time   `json:"updated_at"`
	Priority           int         `json:"priority"`
	NumTimesRewatched  int         `json:"num_times_rewatched"`
	RewatchValue       int         `json:"rewatch_value"`
	Tags               []string    `json:"tags"`
	Comments           string      `json:"comments"`
	StartDate          string      `json:"start_date"`
	FinishDate         string      `json:"finish_date"`
}

// animeList represents the anime list of a user.
type animeList common.ListWithPagination[[]UserAnime]

func (a animeList) Pagination() common.Paging { return a.Paging }

// AnimeList gets the anime list of the user indicated by username (or use @me).
// The anime can be sorted and filtered using the AnimeStatus and SortAnimeList
// option functions respectively.
func (s *UserService) AnimeList(ctx context.Context, username string, options ...AnimeListOption) ([]UserAnime, *Response, error) {
	oo := make([]common.OptionalParam, len(options))
	for i := range options {
		oo[i] = optionFromAnimeListOption(options[i])
	}
	list := new(animeList)
	resp, err := s.client.list(ctx, fmt.Sprintf("users/%s/animelist", username), list, oo...)
	if err != nil {
		return nil, resp, err
	}
	return list.Data, resp, nil
}

func optionFromAnimeListOption(o AnimeListOption) common.OptionFunc {
	return common.OptionFunc(func(v *url.Values) {
		o.AnimeListApply(v)
	})
}

// MARK: interface UpdateMyAnimeListStatusOption
// UpdateMyAnimeListStatusOption are options specific to the
// AnimeService.UpdateMyListStatus method.
type UpdateMyAnimeListStatusOption interface {
	updateMyAnimeListStatusApply(fnSet *url.Values)
}

// MARK: UpdateMyListStatus
// UpdateMyListStatus adds the anime specified by animeID to the user's anime
// list with one or more options added to update the status. If the anime
// already exists in the list, only the status is updated.
func (s *AnimeService) UpdateMyListStatus(ctx context.Context, animeID int, options ...UpdateMyAnimeListStatusOption) (*AnimeListStatus, *Response, error) {
	u := fmt.Sprintf("anime/%d/my_list_status", animeID)
	rawOptions := make([]func(v *url.Values), len(options))
	for i := range options {
		rawOptions[i] = rawOptionFromUpdateMyAnimeListStatusOption(options[i])
	}
	req, err := s.client.NewRequest(http.MethodPatch, u, rawOptions...)
	if err != nil {
		return nil, nil, err
	}

	a := new(AnimeListStatus)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

func rawOptionFromUpdateMyAnimeListStatusOption(o UpdateMyAnimeListStatusOption) func(v *url.Values) {
	return func(v *url.Values) {
		o.updateMyAnimeListStatusApply(v)
	}
}

// MARK: DeleteMyListItem
// DeleteMyListItem deletes an anime from the user's list. If the anime does not
// exist in the user's list, 404 Not Found error is returned.
func (s *AnimeService) DeleteMyListItem(ctx context.Context, animeID int) (*Response, error) {
	u := fmt.Sprintf("anime/%d/my_list_status", animeID)
	req, err := s.client.NewRequest(http.MethodDelete, u)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
