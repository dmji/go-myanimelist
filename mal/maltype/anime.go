package maltype

import (
	"time"

	"github.com/dmji/go-myanimelist/mal/prm"
)

// Anime represents a MyAnimeList anime.
type Anime struct {
	ID                     int                `json:"id"`
	Title                  string             `json:"title"`
	MainPicture            Picture            `json:"main_picture"`
	AlternativeTitles      Titles             `json:"alternative_titles"`
	StartDate              string             `json:"start_date"`
	EndDate                string             `json:"end_date"`
	Synopsis               string             `json:"synopsis"`
	Mean                   float64            `json:"mean"`
	Rank                   int                `json:"rank"`
	Popularity             int                `json:"popularity"`
	NumListUsers           int                `json:"num_list_users"`
	NumScoringUsers        int                `json:"num_scoring_users"`
	NSFW                   string             `json:"nsfw"`
	CreatedAt              time.Time          `json:"created_at"`
	UpdatedAt              time.Time          `json:"updated_at"`
	MediaType              string             `json:"media_type"`
	Status                 string             `json:"status"`
	Genres                 []Genre            `json:"genres"`
	MyListStatus           AnimeListStatus    `json:"my_list_status"`
	NumEpisodes            int                `json:"num_episodes"`
	StartSeason            StartSeason        `json:"start_season"`
	Broadcast              Broadcast          `json:"broadcast"`
	Source                 string             `json:"source"`
	AverageEpisodeDuration int                `json:"average_episode_duration"`
	Rating                 string             `json:"rating"`
	Pictures               []Picture          `json:"pictures"`
	Background             string             `json:"background"`
	RelatedAnime           []RelatedAnime     `json:"related_anime"`
	RelatedManga           []RelatedManga     `json:"related_manga"`
	Recommendations        []RecommendedAnime `json:"recommendations"`
	Studios                []Studio           `json:"studios"`
	Statistics             Statistics         `json:"statistics"`
}

// AnimeListStatus shows the status of each anime in a user's anime list.
type AnimeListStatus struct {
	Status             prm.AnimeStatus `json:"status"`
	Score              int             `json:"score"`
	NumEpisodesWatched int             `json:"num_episodes_watched"`
	IsRewatching       bool            `json:"is_rewatching"`
	UpdatedAt          time.Time       `json:"updated_at"`
	Priority           int             `json:"priority"`
	NumTimesRewatched  int             `json:"num_times_rewatched"`
	RewatchValue       int             `json:"rewatch_value"`
	Tags               []string        `json:"tags"`
	Comments           string          `json:"comments"`
	StartDate          string          `json:"start_date"`
	FinishDate         string          `json:"finish_date"`
}

// Statistics about the anime.
type Statistics struct {
	Status       Status `json:"status"`
	NumListUsers int    `json:"num_list_users"`
}

// StartSeason is the season an anime starts.
type StartSeason struct {
	Year   int    `json:"year"`
	Season string `json:"season"`
}

// Broadcast is the day and time that the show broadcasts.
type Broadcast struct {
	DayOfTheWeek string `json:"day_of_the_week"`
	StartTime    string `json:"start_time"`
}

// RecommendedAnime is a recommended anime.
type RecommendedAnime struct {
	Node               Anime `json:"node"`
	NumRecommendations int   `json:"num_recommendations"`
}

// The Studio that created the anime.
type Studio struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Status of the user's anime list contained in statistics.
type Status struct {
	Watching    string `json:"watching"`
	Completed   string `json:"completed"`
	OnHold      string `json:"on_hold"`
	Dropped     string `json:"dropped"`
	PlanToWatch string `json:"plan_to_watch"`
}

// UserAnime contains an anime record along with its status on the user's list.
type UserAnime struct {
	Anime  Anime           `json:"node"`
	Status AnimeListStatus `json:"list_status"`
}
