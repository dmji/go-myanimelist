package mal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/malhttp"
	"github.com/dmji/go-myanimelist/mal/maltype"
	"github.com/dmji/go-myanimelist/mal/prm"
)

// ForumService handles communication with the forum related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/forum
type ForumService struct {
	client *malhttp.Client

	TopicsOptions       prm.TopicsOptionProvider
	TopicDetailsOptions prm.PagingOptionProvider
}

// NewForumService returns a new ForumService.
func NewForumService(client *malhttp.Client) *ForumService {
	return &ForumService{
		client: client,
	}
}

const (
	boardsEndpoint = "forum/boards"
	topicsEndpoint = "forum/topics"
	topicEndpoint  = "forum/topic"
)

// Boards returns the forum boards.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/forum_boards_get
func (s *ForumService) Boards(ctx context.Context) (*maltype.Forum, *malhttp.Response, error) {
	f := new(maltype.Forum)
	resp, err := s.client.RequestGet(ctx, boardsEndpoint, f)
	if err != nil {
		return nil, resp, err
	}
	return f, resp, nil
}

// TopicDetails returns details about the forum topic specified by topicID.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/forum_topic_get
func (s *ForumService) TopicDetails(ctx context.Context, topicID int, options ...prm.PagingOption) (maltype.TopicDetails, *malhttp.Response, error) {
	rawOptions := optionsToFuncs(options, func(t prm.PagingOption) func(*url.Values) { return t.PagingApply })
	topicDetails, resp, err := s.client.RequestTopicDetails(ctx, fmt.Sprintf("%s/%d", topicEndpoint, topicID), rawOptions...)
	if err != nil {
		return maltype.TopicDetails{}, resp, err
	}
	return topicDetails, resp, nil
}

// Topics returns the forum's topics. Make sure to pass at least the Query
// option or you will get an API error.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/forum_topics_get
func (s *ForumService) Topics(ctx context.Context, options ...prm.TopicsOption) ([]maltype.Topic, *malhttp.Response, error) {
	rawOptions := optionsToFuncs(options, func(t prm.TopicsOption) func(*url.Values) { return t.TopicsApply })
	topics, resp, err := s.client.RequestTopics(ctx, topicsEndpoint, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return topics, resp, nil
}
