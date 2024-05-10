package prm

import (
	"net/url"
	"strings"
)

// Fields is an option that allows to choose the fields that should be returned
// as by default, the API doesn't return all fields.
//
// Example:
//
//	Fields{"synopsis", "my_list_status{priority,comments}"}
type Fields []string

func (f Fields) SeasonalAnimeApply(v *url.Values) { f.Apply(v) }
func (f Fields) AnimeListApply(v *url.Values)     { f.Apply(v) }
func (f Fields) MangaListApply(v *url.Values)     { f.Apply(v) }
func (f Fields) DetailsApply(v *url.Values)       { f.Apply(v) }
func (f Fields) MyInfoApply(v *url.Values)        { f.Apply(v) }
func (f Fields) Apply(v *url.Values) {
	if len(f) != 0 {
		v.Set("fields", strings.Join(f, ","))
	}
}

type AnimeFields struct {
	ID                     bool // "id"
	Title                  bool // "title"
	MainPicture            bool // "main_picture"
	AlternativeTitles      bool // "alternative_titles"
	StartDate              bool // "start_date"
	EndDate                bool // "end_date"
	Synopsis               bool // "synopsis"
	Mean                   bool // "mean"
	Rank                   bool // "rank"
	Popularity             bool // "popularity"
	NumListUsers           bool // "num_list_users"
	NumScoringUsers        bool // "num_scoring_users"
	NSFW                   bool // "nsfw"
	CreatedAt              bool // "created_at"
	UpdatedAt              bool // "updated_at"
	MediaType              bool // "media_type"
	Status                 bool // "status"
	Genres                 bool // "genres"
	MyListStatus           bool // "my_list_status"
	NumEpisodes            bool // "num_episodes"
	StartSeason            bool // "start_season"
	Broadcast              bool // "broadcast"
	Source                 bool // "source"
	AverageEpisodeDuration bool // "average_episode_duration"
	Rating                 bool // "rating"
	Pictures               bool // "pictures"
	Background             bool // "background"
	RelatedAnime           bool // "related_anime"
	RelatedManga           bool // "related_manga"
	Recommendations        bool // "recommendations"
	Studios                bool // "studios"
	Statistics             bool // "statistics"
}

func appendIf(b bool, s []string, v string) {
	if b {
		s = append(s, v)
	}
}

func (f AnimeFields) Fields() Fields {
	res := make([]string, 0, 32)

	appendIf(f.ID, res, "id")
	appendIf(f.Title, res, "title")
	appendIf(f.MainPicture, res, "main_picture")
	appendIf(f.AlternativeTitles, res, "alternative_titles")
	appendIf(f.StartDate, res, "start_date")
	appendIf(f.EndDate, res, "end_date")
	appendIf(f.Synopsis, res, "synopsis")
	appendIf(f.Mean, res, "mean")
	appendIf(f.Rank, res, "rank")
	appendIf(f.Popularity, res, "popularity")
	appendIf(f.NumListUsers, res, "num_list_users")
	appendIf(f.NumScoringUsers, res, "num_scoring_users")
	appendIf(f.NSFW, res, "nsfw")
	appendIf(f.CreatedAt, res, "created_at")
	appendIf(f.UpdatedAt, res, "updated_at")
	appendIf(f.MediaType, res, "media_type")
	appendIf(f.Status, res, "status")
	appendIf(f.Genres, res, "genres")
	appendIf(f.MyListStatus, res, "my_list_status")
	appendIf(f.NumEpisodes, res, "num_episodes")
	appendIf(f.StartSeason, res, "start_season")
	appendIf(f.Broadcast, res, "broadcast")
	appendIf(f.Source, res, "source")
	appendIf(f.AverageEpisodeDuration, res, "average_episode_duration")
	appendIf(f.Rating, res, "rating")
	appendIf(f.Pictures, res, "pictures")
	appendIf(f.Background, res, "background")
	appendIf(f.RelatedAnime, res, "related_anime")
	appendIf(f.RelatedManga, res, "related_manga")
	appendIf(f.Recommendations, res, "recommendations")
	appendIf(f.Studios, res, "studios")
	appendIf(f.Statistics, res, "statistics")

	return res
}

func (f Fields) f(v ...string) Fields {
	res := make([]string, 0, len(v))
	res = append(res, v...)
	return Fields(res)
}
