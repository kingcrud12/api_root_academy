package out

type IPasswordHasher interface {
	hashpassword(password string) string
}
