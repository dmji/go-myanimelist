package mal_prm

type ForumTopicDetailsRequestParameters struct {
	Limit  int `qs:"limit,omitempty"`
	Offset int `qs:"offset,omitempty"`
}
