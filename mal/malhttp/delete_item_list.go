package malhttp

import (
	"context"
	"fmt"
	"net/http"
)

func (s *Client) DeleteMyListItem(ctx context.Context, path string, animeID int) (*Response, error) {
	u := fmt.Sprintf("%s/%d/my_list_status", path, animeID)
	req, err := s.NewRequest(http.MethodDelete, u)
	if err != nil {
		return nil, err
	}

	resp, err := s.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
