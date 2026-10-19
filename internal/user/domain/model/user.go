package model

import "github.com/google/uuid"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Roles     []Role
	Email     string
	Password  string

	InstructorTitle string
	InstructorBio   string

	Billing *BillingInfo

	Avatar string
}

type BillingInfo struct {
	StripeAccountID          string
	StripeCustomerID         string
	StripeOnboardingComplete bool
}

func CreateUser(FirsName string, LastName string, Email string, Password string) *User {
	return &User{
		ID:        uuid.NewString(),
		FirstName: FirsName,
		LastName:  LastName,
		Email:     Email,
		Password:  Password,
	}
}
