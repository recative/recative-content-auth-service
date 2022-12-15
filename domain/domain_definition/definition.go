package domain_definition

type JwtPayload struct {
	UserId      string
	Permissions []string
}
