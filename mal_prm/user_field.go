package mal_prm

//go:generate go run github.com/dmji/go-stringer@latest -type=UserFieldType -trimprefix=UserFieldType -output user_field_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower

import (
	qs "github.com/dmji/qs"
)

type UserFieldType uint8

const (
	UserFieldTypeID              UserFieldType = iota // "id" + argJoin(p...)
	UserFieldTypeName                                 // "name" + argJoin(p...)
	UserFieldTypePicture                              // "picture" + argJoin(p...)
	UserFieldTypeGender                               // "gender" + argJoin(p...)
	UserFieldTypeBirthday                             // "birthday" + argJoin(p...)
	UserFieldTypeLocation                             // "location" + argJoin(p...)
	UserFieldTypeJoinedAt                             // "joined_at" + argJoin(p...)
	UserFieldTypeAnimeStatistics                      // "anime_statistics" + argJoin(p...)
	UserFieldTypeTimeZone                             // "time_zone" + argJoin(p...)
	UserFieldTypeIsSupporter                          // "is_supporter" + argJoin(p...)
)

func (f UserFieldType) UserField(args ...string) UserField {
	return UserField{
		Field: f,
		Args:  args,
	}
}

type UserField struct {
	Field UserFieldType
	Args  []string
}

func (e UserField) MarshalQS(opts *qs.MarshalOptions) ([]string, error) {
	return []string{e.Field.String() + argJoin(e.Args...)}, nil
}
