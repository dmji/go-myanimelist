package mal_opt

// AnimeSeason is the airing season of the anime.
type AnimeSeason string

const (
	// AnimeSeasonWinter is the winter season of January, February and March.
	AnimeSeasonWinter AnimeSeason = "winter"
	// AnimeSeasonSpring is the spring season of April, May and June.
	AnimeSeasonSpring AnimeSeason = "spring"
	// AnimeSeasonSummer is the summer season of July, August and September.
	AnimeSeasonSummer AnimeSeason = "summer"
	// AnimeSeasonFall is the fall season of October, November and December.
	AnimeSeasonFall AnimeSeason = "fall"
)

func (s *AnimeSeason) Winter() AnimeSeason { return AnimeSeasonWinter }
func (s *AnimeSeason) Spring() AnimeSeason { return AnimeSeasonSpring }
func (s *AnimeSeason) Summer() AnimeSeason { return AnimeSeasonSummer }
func (s *AnimeSeason) Fall() AnimeSeason   { return AnimeSeasonFall }
