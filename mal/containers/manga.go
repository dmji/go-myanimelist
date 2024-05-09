package containers

import (
	"time"

	"github.com/dmji/go-myanimelist/mal/prm"
)

// Manga represents a MyAnimeList manga.
type Manga struct {
	ID                int                `json:"id"`
	Title             string             `json:"title"`
	MainPicture       Picture            `json:"main_picture"`
	AlternativeTitles Titles             `json:"alternative_titles"`
	StartDate         string             `json:"start_date"`
	Synopsis          string             `json:"synopsis"`
	Mean              float64            `json:"mean"`
	Rank              int                `json:"rank"`
	Popularity        int                `json:"popularity"`
	NumListUsers      int                `json:"num_list_users"`
	NumScoringUsers   int                `json:"num_scoring_users"`
	Nsfw              string             `json:"nsfw"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	MediaType         string             `json:"media_type"`
	Status            string             `json:"status"`
	Genres            []Genre            `json:"genres"`
	MyListStatus      MangaListStatus    `json:"my_list_status"`
	NumVolumes        int                `json:"num_volumes"`
	NumChapters       int                `json:"num_chapters"`
	Authors           []Author           `json:"authors"`
	Pictures          []Picture          `json:"pictures"`
	Background        string             `json:"background"`
	RelatedAnime      []RelatedAnime     `json:"related_anime"`
	RelatedManga      []RelatedManga     `json:"related_manga"`
	Recommendations   []RecommendedManga `json:"recommendations"`
	Serialization     []Serialization    `json:"serialization"`
}

// Person is usually the creator of a manga.
type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Author of the manga.
type Author struct {
	Person Person `json:"node"`
	Role   string `json:"role"`
}

// RecommendedManga is a manga recommended to the user.
type RecommendedManga struct {
	Node               Manga `json:"node"`
	NumRecommendations int   `json:"num_recommendations"`
}

// Magazine in which the manga was serialized.
type Magazine struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Serialization is a serialized manga series.
type Serialization struct {
	Node Magazine `json:"node"`
	Role string   `json:"role"`
}

// MangaListStatus shows the status of each manga in a user's manga list.
type MangaListStatus struct {
	Status          prm.MangaStatus `json:"status"`
	IsRereading     bool            `json:"is_rereading"`
	NumVolumesRead  int             `json:"num_volumes_read"`
	NumChaptersRead int             `json:"num_chapters_read"`
	Score           int             `json:"score"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Priority        int             `json:"priority"`
	NumTimesReread  int             `json:"num_times_reread"`
	RereadValue     int             `json:"reread_value"`
	Tags            []string        `json:"tags"`
	Comments        string          `json:"comments"`
	StartDate       string          `json:"start_date"`
	FinishDate      string          `json:"finish_date"`
}

// UserManga contains a manga record along with its status on the user's list.
type UserManga struct {
	Manga  Manga           `json:"node"`
	Status MangaListStatus `json:"list_status"`
}
