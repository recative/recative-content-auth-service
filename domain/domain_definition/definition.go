package domain_definition

type JwtPayload struct {
	Permissions []string `json:"permissions"`
}
