package mal

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/dmji/go-myanimelist/mal/common"
)

// AnimeService handles communication with the anime related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/anime
// https://myanimelist.net/apiconfig/references/api/v2#tag/user-animelist
type AnimeService struct {
	client *Client
}

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

// Picture is a representative picture from the show.
type Picture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

// Titles of the anime in English and Japanese.
type Titles struct {
	Synonyms []string `json:"synonyms"`
	En       string   `json:"en"`
	Ja       string   `json:"ja"`
}

// The Genre of the anime.
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

// Statistics about the anime.
type Statistics struct {
	Status       Status `json:"status"`
	NumListUsers int    `json:"num_list_users"`
}

// RecommendedAnime is a recommended anime.
type RecommendedAnime struct {
	Node               Anime `json:"node"`
	NumRecommendations int   `json:"num_recommendations"`
}

// RelatedAnime contains a related anime.
type RelatedAnime struct {
	Node                  Anime  `json:"node"`
	RelationType          string `json:"relation_type"`
	RelationTypeFormatted string `json:"relation_type_formatted"`
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

// DetailsOption is an option specific for the anime and manga details methods.
type DetailsOption interface {
	DetailsApply(v *url.Values)
}

// Details returns details about an anime. By default, few anime fields are
// populated. Use the Fields option to specify which fields should be included.
func (s *AnimeService) Details(ctx context.Context, animeID int, options ...DetailsOption) (*Anime, *Response, error) {
	a := new(Anime)
	resp, err := s.client.details(ctx, fmt.Sprintf("anime/%d", animeID), a, options...)
	if err != nil {
		return nil, resp, err
	}
	return a, resp, nil
}

// List allows an authenticated user to search and list anime data. You may get
// user specific data by using the optional field "my_list_status".
func (s *AnimeService) List(ctx context.Context, search string, options ...common.OptionalParam) ([]Anime, *Response, error) {
	options = append(options, common.OptionFromQuery(search))
	return s.list(ctx, "anime", options...)
}

func (s *AnimeService) list(ctx context.Context, path string, options ...common.OptionalParam) ([]Anime, *Response, error) {
	list := new(animeList)
	resp, err := s.client.list(ctx, path, list, options...)
	if err != nil {
		return nil, resp, err
	}
	anime := make([]Anime, len(list.Data))
	for i := range list.Data {
		anime[i] = list.Data[i].Anime
	}
	return anime, resp, nil
}

// Suggested returns suggested anime for the authorized user. If the user is new
// comer, this endpoint returns an empty list.
func (s *AnimeService) Suggested(ctx context.Context, options ...common.OptionalParam) ([]Anime, *Response, error) {
	return s.list(ctx, "anime/suggestions", options...)
}
