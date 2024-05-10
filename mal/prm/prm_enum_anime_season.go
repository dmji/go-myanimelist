package prm

// AnimeSeason is the airing season of the anime.
type AnimeSeason string

const (
	// AnimeSeasonWinter is the winter season of January, February and March.
	animeSeasonWinter AnimeSeason = "winter"
	// AnimeSeasonSpring is the spring season of April, May and June.
	animeSeasonSpring AnimeSeason = "spring"
	// AnimeSeasonSummer is the summer season of July, August and September.
	animeSeasonSummer AnimeSeason = "summer"
	// AnimeSeasonFall is the fall season of October, November and December.
	animeSeasonFall AnimeSeason = "fall"
)

func (s *AnimeSeason) Winter() AnimeSeason { return animeSeasonWinter }
func (s *AnimeSeason) Spring() AnimeSeason { return animeSeasonSpring }
func (s *AnimeSeason) Summer() AnimeSeason { return animeSeasonSummer }
func (s *AnimeSeason) Fall() AnimeSeason   { return animeSeasonFall }
