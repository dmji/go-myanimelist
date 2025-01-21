package mal_prm

type UserMyInfoRequestParameters struct {
	Fields []UserField `qs:"fields,omitempty"`
}
