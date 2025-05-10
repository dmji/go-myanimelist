package mal_prm

//go:generate go run github.com/dmji/go-stringer@latest -type=MangaRanking,AnimeRanking -trimprefix=@me -output ranking_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower

// MangaRanking allows to choose how the manga will be ranked.
type MangaRanking int8

const (
	// MangaRankingAll returns all types ranked.
	MangaRankingAll MangaRanking = iota // "all"
	// MangaRankingManga returns the top manga.
	MangaRankingManga // "manga"
	// MangaRankingOneshots returns the top one-shots.
	MangaRankingOneshots // "oneshots"
	// MangaRankingDoujinshi returns the top doujinshi.
	MangaRankingDoujin // "doujin"
	// MangaRankingLightNovels returns the top light novels.
	MangaRankingLightNovels // "lightnovels"
	// MangaRankingNovels returns the top novels.
	MangaRankingNovels // "novels"
	// MangaRankingManhwa returns the top manhwa.
	MangaRankingManhwa // "manhwa"
	// MangaRankingManhua returns the top manhua.
	MangaRankingManhua // "manhua"
	// MangaRankingByPopularity returns the most popular entries.
	MangaRankingByPopularity // "bypopularity"
	// MangaRankingFavorite returns the most favorite entries.
	MangaRankingFavorite // "favorite"
)

// AnimeRanking allows to choose how the anime will be ranked.
type AnimeRanking uint8

const (
	// AnimeRankingAll returns the top anime series.
	AnimeRankingAll AnimeRanking = iota // "all"
	// AnimeRankingAiring returns the top airing anime.
	AnimeRankingAiring // "airing"
	// AnimeRankingUpcoming returns the top upcoming anime.
	AnimeRankingUpcoming // "upcoming"
	// AnimeRankingTV returns the top Anime TV series.
	AnimeRankingTV // "tv"
	// AnimeRankingOVA returns the top anime OVA series.
	AnimeRankingOVA // "ova"
	// AnimeRankingMovie returns the top anime movies.
	AnimeRankingMovie // "movie"
	// AnimeRankingSpecial returns the top anime specials.
	AnimeRankingSpecial // "special"
	// AnimeRankingByPopularity returns the top anime by popularity.
	AnimeRankingByPopularity // "bypopularity"
	// AnimeRankingFavorite returns the top favorite Anime.
	AnimeRankingFavorite // "favorite"
)
