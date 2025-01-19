package mal_opt

import "net/url"

// MyInfoOption are options specific to the User.MyInfo method.
type MyInfoOption interface {
	MyInfoApply(v *url.Values)
}

type MyInfoOptionProvider struct {
	UserFields
}

func (s MyInfoOptionProvider) Fields(v ...string) Fields {
	return NewFields(v...)
}
