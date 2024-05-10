package prm

import (
	"net/url"
)

// BoardID is an option that filters topics based on the board ID.
type BoardID int

func (id BoardID) TopicsApply(v *url.Values) { v.Set("board_id", itoa(int(id))) }
func (s BoardID) Val(v int) BoardID          { return BoardID(v) }

// SubboardID is an option that filters topics based on the subboard ID.
type SubboardID int

func (id SubboardID) TopicsApply(v *url.Values) { v.Set("subboard_id", itoa(int(id))) }
func (s SubboardID) Val(v int) SubboardID       { return SubboardID(v) }

// sortTopics is an option that sorts the returned topics.
type sortTopics string

// SortTopicsRecent is the default and only sorting value for topics.
const SortTopicsRecent sortTopics = "recent"

func (s sortTopics) TopicsApply(v *url.Values) { v.Set("sort", string(s)) }
func (s sortTopics) Recent() sortTopics        { return SortTopicsRecent }

// Query is an option that allows to search for a term.
type Query string

func (q Query) TopicsApply(v *url.Values) { v.Set("q", string(q)) }
func (s Query) Val(v string) Query        { return Query(v) }

// TopicUserName is an option that filters topics based on the topic username.
type TopicUserName string

func (n TopicUserName) TopicsApply(v *url.Values)  { v.Set("topic_user_name", string(n)) }
func (s TopicUserName) Val(v string) TopicUserName { return TopicUserName(v) }

// UserName is an option that filters topics based on a username.
type UserName string

func (n UserName) TopicsApply(v *url.Values) { v.Set("user_name", string(n)) }
func (s UserName) Val(v string) UserName     { return UserName(v) }
