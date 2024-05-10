package prm

import "net/url"

// DetailsOption is an option specific for the anime and manga details methods.
type DetailsOption interface {
	DetailsApply(v *url.Values)
}

type DetailsOptionProvider struct {
	AnimeFields
	MangaFields
}

func (s DetailsOptionProvider) Fields(v ...string) Fields {
	return NewFields(v...)
}
