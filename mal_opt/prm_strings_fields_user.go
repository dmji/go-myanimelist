package mal_opt

type UserFields struct{}

func (f UserFields) ID(p ...string) string {
	return "id" + argJoin(p...)
}

func (f UserFields) Name(p ...string) string {
	return "name" + argJoin(p...)
}

func (f UserFields) Picture(p ...string) string {
	return "picture" + argJoin(p...)
}

func (f UserFields) Gender(p ...string) string {
	return "gender" + argJoin(p...)
}

func (f UserFields) Birthday(p ...string) string {
	return "birthday" + argJoin(p...)
}

func (f UserFields) Location(p ...string) string {
	return "location" + argJoin(p...)
}

func (f UserFields) JoinedAt(p ...string) string {
	return "joined_at" + argJoin(p...)
}

func (f UserFields) AnimeStatistics(p ...string) string {
	return "anime_statistics" + argJoin(p...)
}

func (f UserFields) TimeZone(p ...string) string {
	return "time_zone" + argJoin(p...)
}

func (f UserFields) IsSupporter(p ...string) string {
	return "is_supporter" + argJoin(p...)
}
