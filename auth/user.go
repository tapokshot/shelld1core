package auth

type AuthUser interface {
	GetLogin() string
	GetRoles() []string
}