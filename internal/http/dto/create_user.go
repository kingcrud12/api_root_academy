package dto

type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserCreated struct {
	FirstName string
	LastName  string
	Email     string
}
