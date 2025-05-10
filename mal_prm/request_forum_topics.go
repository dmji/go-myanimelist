package mal_prm

type ForumTopicsRequestParameters struct {
	BoardID       int        `qs:"board_id,omitempty"`
	SubboardID    int        `qs:"subboard_id,omitempty"`
	Query         string     `qs:"q,omitempty"`
	TopicUserName string     `qs:"topic_user_name,omitempty"`
	UserName      string     `qs:"user_name,omitempty"`
	Limit         int        `qs:"limit,omitempty"`
	Offset        int        `qs:"offset,omitempty"`
	Sort          SortTopics `qs:"sort,omitempty"`
}
