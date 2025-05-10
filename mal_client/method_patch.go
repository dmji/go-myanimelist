package mal_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dmji/go-myanimelist/mal_prm"
)

// UpdateMyListStatus adds the manga specified by mangaID to the user's manga
// list with one or more options added to update the status. If the manga
// already exists in the list, only the status is updated.
func (s *Client) UpdateMyListStatus(ctx context.Context, path string, id int, v interface{}, options *mal_prm.UserMangaListRequestParameters) (*Response, error) {
	u := fmt.Sprintf("%s/%d/my_list_status", path, id)
	req, err := s.NewRequest(http.MethodPatch, u, options...)
	if err != nil {
		return nil, err
	}

	resp, err := s.Do(ctx, req, v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
