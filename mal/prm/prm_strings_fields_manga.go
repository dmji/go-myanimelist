package prm

type MangaFields struct {
	ID                bool // "id"
	Title             bool // "title"
	MainPicture       bool // "main_picture"
	AlternativeTitles bool // "alternative_titles"
	StartDate         bool // "start_date"
	Synopsis          bool // "synopsis"
	Mean              bool // "mean"
	Rank              bool // "rank"
	Popularity        bool // "popularity"
	NumListUsers      bool // "num_list_users"
	NumScoringUsers   bool // "num_scoring_users"
	Nsfw              bool // "nsfw"
	CreatedAt         bool // "created_at"
	UpdatedAt         bool // "updated_at"
	MediaType         bool // "media_type"
	Status            bool // "status"
	Genres            bool // "genres"
	MyListStatus      bool // "my_list_status"
	NumVolumes        bool // "num_volumes"
	NumChapters       bool // "num_chapters"
	Authors           bool // "authors"
	AuthorsL          bool // "authors{last_name}"
	AuthorsF          bool // "authors{first_name}"
	AuthorsLF         bool // "authors{last_name, first_name}"
	Pictures          bool // "pictures"
	Background        bool // "background"
	RelatedAnime      bool // "related_anime"
	RelatedManga      bool // "related_manga"
	Recommendations   bool // "recommendations"
	Serialization     bool // "serialization"
}

func (f MangaFields) Fields() Fields {
	res := make([]string, 0, 32)

	appendIf(f.ID, res, "id")
	appendIf(f.Title, res, "title")
	appendIf(f.MainPicture, res, "main_picture")
	appendIf(f.AlternativeTitles, res, "alternative_titles")
	appendIf(f.StartDate, res, "start_date")
	appendIf(f.Synopsis, res, "synopsis")
	appendIf(f.Mean, res, "mean")
	appendIf(f.Rank, res, "rank")
	appendIf(f.Popularity, res, "popularity")
	appendIf(f.NumListUsers, res, "num_list_users")
	appendIf(f.NumScoringUsers, res, "num_scoring_users")
	appendIf(f.Nsfw, res, "nsfw")
	appendIf(f.CreatedAt, res, "created_at")
	appendIf(f.UpdatedAt, res, "updated_at")
	appendIf(f.MediaType, res, "media_type")
	appendIf(f.Status, res, "status")
	appendIf(f.Genres, res, "genres")
	appendIf(f.MyListStatus, res, "my_list_status")
	appendIf(f.NumVolumes, res, "num_volumes")
	appendIf(f.NumChapters, res, "num_chapters")
	appendIf(f.Authors, res, "authors")
	appendIf(f.AuthorsL, res, "authors{last_name}")
	appendIf(f.AuthorsF, res, "authors{first_name}")
	appendIf(f.AuthorsLF, res, "authors{last_name, first_name}")
	appendIf(f.Pictures, res, "pictures")
	appendIf(f.Background, res, "background")
	appendIf(f.RelatedAnime, res, "related_anime")
	appendIf(f.RelatedManga, res, "related_manga")
	appendIf(f.Recommendations, res, "recommendations")
	appendIf(f.Serialization, res, "serialization")

	return res
}

func (f MangaFields) F(v ...string) Fields {
	res := make([]string, 0, len(v))
	res = append(res, v...)
	return Fields(res)
}
