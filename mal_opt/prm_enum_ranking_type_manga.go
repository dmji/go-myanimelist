package mal_opt

// MangaRanking allows to choose how the manga will be ranked.
type MangaRanking string

const (
	// MangaRankingAll returns all types ranked.
	MangaRankingAll MangaRanking = "all"
	// MangaRankingManga returns the top manga.
	MangaRankingManga MangaRanking = "manga"
	// MangaRankingOneshots returns the top one-shots.
	MangaRankingOneshots MangaRanking = "oneshots"
	// MangaRankingDoujinshi returns the top doujinshi.
	MangaRankingDoujinshi MangaRanking = "doujin"
	// MangaRankingLightNovels returns the top light novels.
	MangaRankingLightNovels MangaRanking = "lightnovels"
	// MangaRankingNovels returns the top novels.
	MangaRankingNovels MangaRanking = "novels"
	// MangaRankingManhwa returns the top manhwa.
	MangaRankingManhwa MangaRanking = "manhwa"
	// MangaRankingManhua returns the top manhua.
	MangaRankingManhua MangaRanking = "manhua"
	// MangaRankingByPopularity returns the most popular entries.
	MangaRankingByPopularity MangaRanking = "bypopularity"
	// MangaRankingFavorite returns the most favorite entries.
	MangaRankingFavorite MangaRanking = "favorite"
)

func (r MangaRanking) All() MangaRanking          { return MangaRankingAll }
func (r MangaRanking) Manga() MangaRanking        { return MangaRankingManga }
func (r MangaRanking) Oneshots() MangaRanking     { return MangaRankingOneshots }
func (r MangaRanking) Doujinshi() MangaRanking    { return MangaRankingDoujinshi }
func (r MangaRanking) LightNovels() MangaRanking  { return MangaRankingLightNovels }
func (r MangaRanking) Novels() MangaRanking       { return MangaRankingNovels }
func (r MangaRanking) Manhwa() MangaRanking       { return MangaRankingManhwa }
func (r MangaRanking) Manhua() MangaRanking       { return MangaRankingManhua }
func (r MangaRanking) ByPopularity() MangaRanking { return MangaRankingByPopularity }
func (r MangaRanking) Favorite() MangaRanking     { return MangaRankingFavorite }
