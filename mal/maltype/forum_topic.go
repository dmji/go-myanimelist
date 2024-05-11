package maltype

import "time"

// A Topic of the forum.
type Topic struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         CreatedBy `json:"created_by"`
	NumberOfPosts     int       `json:"number_of_posts"`
	LastPostCreatedAt time.Time `json:"last_post_created_at"`
	LastPostCreatedBy CreatedBy `json:"last_post_created_by"`
	IsLocked          bool      `json:"is_locked"`
}

// TopicDetails contain the posts of a forum topic and an optional poll.
type TopicDetails struct {
	Title string `json:"title"`
	Posts []Post `json:"posts"`
	Poll  *Poll  `json:"poll"`
}

// Post is a forum post.
type Post struct {
	ID        int       `json:"id"`
	Number    int       `json:"number"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy CreatedBy `json:"created_by"`
	Body      string    `json:"body"`
	Signature string    `json:"signature"`
}

// CreatedBy shows the name of the user that created the post or topic.
type CreatedBy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ForumAvator string `json:"forum_avator"`
}

// Poll is an optional poll in a forum post.
type Poll struct {
	ID       int          `json:"id"`
	Question string       `json:"question"`
	Closed   bool         `json:"closed"`
	Options  []PollOption `json:"options"`
}

// PollOption is one of the choices of a poll.
type PollOption struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Votes int    `json:"votes"`
}
