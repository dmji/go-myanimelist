package prm

// AnimeRanking allows to choose how the anime will be ranked.
type AnimeRanking string

const (
	// AnimeRankingAll returns the top anime series.
	AnimeRankingAll AnimeRanking = "all"
	// AnimeRankingAiring returns the top airing anime.
	AnimeRankingAiring AnimeRanking = "airing"
	// AnimeRankingUpcoming returns the top upcoming anime.
	AnimeRankingUpcoming AnimeRanking = "upcoming"
	// AnimeRankingTV returns the top Anime TV series.
	AnimeRankingTV AnimeRanking = "tv"
	// AnimeRankingOVA returns the top anime OVA series.
	AnimeRankingOVA AnimeRanking = "ova"
	// AnimeRankingMovie returns the top anime movies.
	AnimeRankingMovie AnimeRanking = "movie"
	// AnimeRankingSpecial returns the top anime specials.
	AnimeRankingSpecial AnimeRanking = "special"
	// AnimeRankingByPopularity returns the top anime by popularity.
	AnimeRankingByPopularity AnimeRanking = "bypopularity"
	// AnimeRankingFavorite returns the top favorite Anime.
	AnimeRankingFavorite AnimeRanking = "favorite"
)
