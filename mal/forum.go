package mal

import (
	"context"
	"fmt"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_prm"
	"github.com/dmji/go-myanimelist/mal_type"
)

type clientForum interface {
	RequestGet(ctx context.Context, path string, v interface{}, q interface{}) (*mal_client.Response, error)
	RequestTopicDetails(ctx context.Context, path string, opts *mal_prm.ForumTopicDetailsRequestParameters) (mal_type.TopicDetails, *mal_client.Response, error)
	RequestTopics(ctx context.Context, path string, opts *mal_prm.ForumTopicsRequestParameters) ([]mal_type.Topic, *mal_client.Response, error)
}

// ForumService handles communication with the forum related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/forum
type ForumService struct {
	client clientForum
}

// NewForumService returns a new ForumService.
func NewForumService(client clientForum) *ForumService {
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
func (s *ForumService) Boards(ctx context.Context) (*mal_type.Forum, *mal_client.Response, error) {
	f := new(mal_type.Forum)
	resp, err := s.client.RequestGet(ctx, boardsEndpoint, f, nil)
	if err != nil {
		return nil, resp, err
	}
	return f, resp, nil
}

// TopicDetails returns details about the forum topic specified by topicID.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/forum_topic_get
func (s *ForumService) TopicDetails(ctx context.Context, topicID int, opts *mal_prm.ForumTopicDetailsRequestParameters) (mal_type.TopicDetails, *mal_client.Response, error) {
	if opts == nil {
		opts = &mal_prm.ForumTopicDetailsRequestParameters{}
	}

	topicDetails, resp, err := s.client.RequestTopicDetails(ctx, fmt.Sprintf("%s/%d", topicEndpoint, topicID), opts)
	if err != nil {
		return mal_type.TopicDetails{}, resp, err
	}
	return topicDetails, resp, nil
}

// Topics returns the forum's topics. Make sure to pass at least the Query
// option or you will get an API error.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/forum_topics_get
func (s *ForumService) Topics(ctx context.Context, opts *mal_prm.ForumTopicsRequestParameters) ([]mal_type.Topic, *mal_client.Response, error) {
	if opts == nil {
		opts = &mal_prm.ForumTopicsRequestParameters{}
	}

	topics, resp, err := s.client.RequestTopics(ctx, topicsEndpoint, opts)
	if err != nil {
		return nil, resp, err
	}
	return topics, resp, nil
}
