package mal_opt

type AnimeFields struct{}

func (f AnimeFields) ID(p ...string) string {
	return "id" + argJoin(p...)
}

func (f AnimeFields) Title(p ...string) string {
	return "title" + argJoin(p...)
}

func (f AnimeFields) MainPicture(p ...string) string {
	return "main_picture" + argJoin(p...)
}

func (f AnimeFields) AlternativeTitles(p ...string) string {
	return "alternative_titles" + argJoin(p...)
}

func (f AnimeFields) StartDate(p ...string) string {
	return "start_date" + argJoin(p...)
}

func (f AnimeFields) EndDate(p ...string) string {
	return "end_date" + argJoin(p...)
}

func (f AnimeFields) Synopsis(p ...string) string {
	return "synopsis" + argJoin(p...)
}

func (f AnimeFields) Mean(p ...string) string {
	return "mean" + argJoin(p...)
}

func (f AnimeFields) Rank(p ...string) string {
	return "rank" + argJoin(p...)
}

func (f AnimeFields) Popularity(p ...string) string {
	return "popularity" + argJoin(p...)
}

func (f AnimeFields) NumListUsers(p ...string) string {
	return "num_list_users" + argJoin(p...)
}

func (f AnimeFields) NumScoringUsers(p ...string) string {
	return "num_scoring_users" + argJoin(p...)
}

func (f AnimeFields) NSFW(p ...string) string {
	return "nsfw" + argJoin(p...)
}

func (f AnimeFields) CreatedAt(p ...string) string {
	return "created_at" + argJoin(p...)
}

func (f AnimeFields) UpdatedAt(p ...string) string {
	return "updated_at" + argJoin(p...)
}

func (f AnimeFields) MediaType(p ...string) string {
	return "media_type" + argJoin(p...)
}

func (f AnimeFields) Status(p ...string) string {
	return "status" + argJoin(p...)
}

func (f AnimeFields) Genres(p ...string) string {
	return "genres" + argJoin(p...)
}

func (f AnimeFields) MyListStatus(p ...string) string {
	return "my_list_status" + argJoin(p...)
}

func (f AnimeFields) NumEpisodes(p ...string) string {
	return "num_episodes" + argJoin(p...)
}

func (f AnimeFields) StartSeason(p ...string) string {
	return "start_season" + argJoin(p...)
}

func (f AnimeFields) Broadcast(p ...string) string {
	return "broadcast" + argJoin(p...)
}

func (f AnimeFields) Source(p ...string) string {
	return "source" + argJoin(p...)
}

func (f AnimeFields) AverageEpisodeDuration(p ...string) string {
	return "average_episode_duration" + argJoin(p...)
}

func (f AnimeFields) Rating(p ...string) string {
	return "rating" + argJoin(p...)
}

func (f AnimeFields) Pictures(p ...string) string {
	return "pictures" + argJoin(p...)
}

func (f AnimeFields) Background(p ...string) string {
	return "background" + argJoin(p...)
}

func (f AnimeFields) RelatedAnime(p ...string) string {
	return "related_anime" + argJoin(p...)
}

func (f AnimeFields) RelatedManga(p ...string) string {
	return "related_manga" + argJoin(p...)
}

func (f AnimeFields) Recommendations(p ...string) string {
	return "recommendations" + argJoin(p...)
}

func (f AnimeFields) Studios(p ...string) string {
	return "studios" + argJoin(p...)
}

func (f AnimeFields) Statistics(p ...string) string {
	return "statistics" + argJoin(p...)
}
