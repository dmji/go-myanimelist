package mal_prm

type UserAnimeListRequestParameters struct {
	Fields []AnimeField    `qs:"fields,omitempty"`
	Limit  int             `qs:"limit,omitempty"`
	Offset int             `qs:"offset,omitempty"`
	NSFW   bool            `qs:"nsfw,omitempty"`
	Status AnimeStatus     `qs:"status,omitempty"`
	Sort   SortAnimeListBy `qs:"sort,omitempty"`
}
