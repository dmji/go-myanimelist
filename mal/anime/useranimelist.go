package anime

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/common"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

// MARK: UpdateMyListStatus
// UpdateMyListStatus adds the anime specified by animeID to the user's anime
// list with one or more options added to update the status. If the anime
// already exists in the list, only the status is updated.
func (s *Service) UpdateMyListStatus(ctx context.Context, animeID int, options ...prm.UpdateMyAnimeListStatusOption) (*containers.AnimeListStatus, *common.Response, error) {
	rawOptions := common.OptionsToFuncs(options, func(t prm.UpdateMyAnimeListStatusOption) func(*url.Values) { return t.UpdateMyAnimeListStatusApply })

	a := new(containers.AnimeListStatus)
	resp, err := s.client.UpdateMyListStatus(ctx, "anime", animeID, a, rawOptions...)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}
