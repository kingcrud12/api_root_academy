package user

type Role string

const (
	RoleTeacher Role = "teacher"
	RoleStudent Role = "student"
	RoleAdmin   Role = "admin"
)

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
