package out

type IPasswordHasher interface {
	HashPassword(password string) (string, error)
}
