package mal_prm

type AnimeDetailsRequestParameters struct {
	Fields []AnimeField `qs:"fields,omitempty"`
}
