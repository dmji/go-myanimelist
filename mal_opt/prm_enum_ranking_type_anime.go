package mal_opt

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

func (r AnimeRanking) All() AnimeRanking          { return AnimeRankingAll }
func (r AnimeRanking) Airing() AnimeRanking       { return AnimeRankingAiring }
func (r AnimeRanking) Upcoming() AnimeRanking     { return AnimeRankingUpcoming }
func (r AnimeRanking) TV() AnimeRanking           { return AnimeRankingTV }
func (r AnimeRanking) OVA() AnimeRanking          { return AnimeRankingOVA }
func (r AnimeRanking) Movie() AnimeRanking        { return AnimeRankingMovie }
func (r AnimeRanking) Special() AnimeRanking      { return AnimeRankingSpecial }
func (r AnimeRanking) ByPopularity() AnimeRanking { return AnimeRankingByPopularity }
func (r AnimeRanking) Favorite() AnimeRanking     { return AnimeRankingFavorite }
