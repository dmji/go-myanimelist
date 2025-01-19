package mal_opt

import (
	"net/url"
)

type TopicsOption interface {
	TopicsApply(v *url.Values)
}

type TopicsOptionProvider struct {
	SortTopics sortTopics
}

func (s TopicsOptionProvider) BoardID(v int) BoardID                { return BoardID(v) }
func (s TopicsOptionProvider) SubboardID(v int) SubboardID          { return SubboardID(v) }
func (s TopicsOptionProvider) Query(v string) Query                 { return Query(v) }
func (s TopicsOptionProvider) TopicUserName(v string) TopicUserName { return TopicUserName(v) }
func (s TopicsOptionProvider) UserName(v string) UserName           { return UserName(v) }
func (s TopicsOptionProvider) Limit(v int) Limit                    { return NewLimit(v) }
func (s TopicsOptionProvider) Offset(v int) Offset                  { return Offset(v) }
