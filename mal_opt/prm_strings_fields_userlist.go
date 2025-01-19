package mal_opt

type UserListFields struct{}

func (f UserListFields) Node(p ...string) string {
	return "node" + argJoin(p...)
}

func (f UserListFields) ListStatus(p ...string) string {
	return "list_status" + argJoin(p...)
}
