package manga

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/common"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

// UpdateMyListStatus adds the manga specified by mangaID to the user's manga
// list with one or more options added to update the status. If the manga
// already exists in the list, only the status is updated.
func (s *Service) UpdateMyListStatus(ctx context.Context, mangaID int, options ...prm.UpdateMyMangaListStatusOption) (*containers.MangaListStatus, *common.Response, error) {
	rawOptions := common.OptionsToFuncs(options, func(t prm.UpdateMyMangaListStatusOption) func(*url.Values) { return t.UpdateMyMangaListStatusApply })

	m := new(containers.MangaListStatus)
	resp, err := s.client.UpdateMyListStatus(ctx, "manga", mangaID, m, rawOptions...)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}
