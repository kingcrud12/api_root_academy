package model

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
