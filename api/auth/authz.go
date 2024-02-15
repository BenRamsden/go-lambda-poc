package auth

type Permission string

const (
	PermissionRead   Permission = "read"
	PermissionWrite  Permission = "write"
	PermissionDelete Permission = "delete"
)

type AuthGroup string

const (
	AuthGroupAdmin AuthGroup = "admin"
	AuthGroupUser  AuthGroup = "user"
)

type AuthService interface {
	// Add or remove permissions for a user
	Enforce(actor string, object string, permission Permission) error
	Revoke(actor string, object string, permission Permission) error
	Can(actor string, object string, permission Permission) bool
}
