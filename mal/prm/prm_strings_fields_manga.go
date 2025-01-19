package prm

type MangaFields struct{}

func (f MangaFields) ID(p ...string) string {
	return "id" + argJoin(p...)
}

func (f MangaFields) Title(p ...string) string {
	return "title" + argJoin(p...)
}

func (f MangaFields) MainPicture(p ...string) string {
	return "main_picture" + argJoin(p...)
}

func (f MangaFields) AlternativeTitles(p ...string) string {
	return "alternative_titles" + argJoin(p...)
}

func (f MangaFields) StartDate(p ...string) string {
	return "start_date" + argJoin(p...)
}

func (f MangaFields) Synopsis(p ...string) string {
	return "synopsis" + argJoin(p...)
}

func (f MangaFields) Mean(p ...string) string {
	return "mean" + argJoin(p...)
}

func (f MangaFields) Rank(p ...string) string {
	return "rank" + argJoin(p...)
}

func (f MangaFields) Popularity(p ...string) string {
	return "popularity" + argJoin(p...)
}

func (f MangaFields) NumListUsers(p ...string) string {
	return "num_list_users" + argJoin(p...)
}

func (f MangaFields) NumScoringUsers(p ...string) string {
	return "num_scoring_users" + argJoin(p...)
}

func (f MangaFields) Nsfw(p ...string) string {
	return "nsfw" + argJoin(p...)
}

func (f MangaFields) CreatedAt(p ...string) string {
	return "created_at" + argJoin(p...)
}

func (f MangaFields) UpdatedAt(p ...string) string {
	return "updated_at" + argJoin(p...)
}

func (f MangaFields) MediaType(p ...string) string {
	return "media_type" + argJoin(p...)
}

func (f MangaFields) Status(p ...string) string {
	return "status" + argJoin(p...)
}

func (f MangaFields) Genres(p ...string) string {
	return "genres" + argJoin(p...)
}

func (f MangaFields) MyListStatus(p ...string) string {
	return "my_list_status" + argJoin(p...)
}

func (f MangaFields) NumVolumes(p ...string) string {
	return "num_volumes" + argJoin(p...)
}

func (f MangaFields) NumChapters(p ...string) string {
	return "num_chapters" + argJoin(p...)
}

func (f MangaFields) Authors(p ...string) string {
	return "authors" + argJoin(p...)
}

func (f MangaFields) Pictures(p ...string) string {
	return "pictures" + argJoin(p...)
}

func (f MangaFields) Background(p ...string) string {
	return "background" + argJoin(p...)
}

func (f MangaFields) RelatedAnime(p ...string) string {
	return "related_anime" + argJoin(p...)
}

func (f MangaFields) RelatedManga(p ...string) string {
	return "related_manga" + argJoin(p...)
}

func (f MangaFields) Recommendations(p ...string) string {
	return "recommendations" + argJoin(p...)
}

func (f MangaFields) Serialization(p ...string) string {
	return "serialization" + argJoin(p...)
}
