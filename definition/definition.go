package definition

type JwtPayload struct {
	UserId      string
	Permissions []string
}

type CrossServiceConfig struct {
	AdminAuthorizationToken string
}
