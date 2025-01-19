package mal_type

// Picture is a representative picture from the show.
type Picture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

// Titles of the anime in English and Japanese.
type Titles struct {
	Synonyms []string `json:"synonyms"`
	En       string   `json:"en"`
	Ja       string   `json:"ja"`
}

// The Genre of the anime.
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RelatedManga shows manga related with the returned entry.
type RelatedManga struct {
	Node                  Manga  `json:"node"`
	RelationType          string `json:"relation_type"`
	RelationTypeFormatted string `json:"relation_type_formatted"`
}

// RelatedAnime contains a related anime.
type RelatedAnime struct {
	Node                  Anime  `json:"node"`
	RelationType          string `json:"relation_type"`
	RelationTypeFormatted string `json:"relation_type_formatted"`
}
