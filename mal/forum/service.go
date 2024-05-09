package forum

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/api_driver"
	"github.com/dmji/go-myanimelist/mal/common"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

// Service handles communication with the forum related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/forum
type Service struct {
	client *api_driver.Client
}

func NewService(client *api_driver.Client) *Service {
	return &Service{
		client: client,
	}
}

const (
	boardsEndpoint = "forum/boards"
	topicsEndpoint = "forum/topics"
	topicEndpoint  = "forum/topic"
)

// Boards returns the forum boards.
func (s *Service) Boards(ctx context.Context) (*containers.Forum, *common.Response, error) {
	f := new(containers.Forum)
	resp, err := s.client.RequestGet(ctx, boardsEndpoint, f)
	if err != nil {
		return nil, resp, err
	}
	return f, resp, nil
}

// Topics returns the forum's topics. Make sure to pass at least the Query
// option or you will get an API error.
func (s *Service) Topics(ctx context.Context, options ...prm.TopicsOption) ([]containers.Topic, *common.Response, error) {
	rawOptions := common.OptionsToFuncs(options, func(t prm.TopicsOption) func(*url.Values) { return t.TopicsApply })
	topics, resp, err := s.client.RequestTopics(ctx, topicsEndpoint, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return topics, resp, nil
}

// TopicDetails returns details about the forum topic specified by topicID.
func (s *Service) TopicDetails(ctx context.Context, topicID int, options ...prm.PagingOption) (containers.TopicDetails, *common.Response, error) {
	rawOptions := common.OptionsToFuncs(options, func(t prm.PagingOption) func(*url.Values) { return t.PagingApply })
	topicDetails, resp, err := s.client.RequestTopicDetails(ctx, fmt.Sprintf("%s/%d", topicEndpoint, topicID), rawOptions...)
	if err != nil {
		return containers.TopicDetails{}, resp, err
	}
	return topicDetails, resp, nil
}
