package mal_prm

type UserMangaListRequestParameters struct {
	Fields []MangaField    `qs:"fields,omitempty"`
	Limit  int             `qs:"limit,omitempty"`
	Offset int             `qs:"offset,omitempty"`
	NSFW   bool            `qs:"nsfw,omitempty"`
	Status MangaStatus     `qs:"status,omitempty"`
	Sort   SortMangaListBy `qs:"sort,omitempty"`
}
