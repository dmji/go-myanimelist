package mal_prm

type MangaDetailsRequestParameters struct {
	Fields []MangaField `qs:"fields,omitempty"`
}
