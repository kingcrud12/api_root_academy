package postgres

import (
	"strings"

	"root_academy_go/internal/domain/model/user"
)

// UserEntity is the GORM persistence entity for user.User. It stays in the
// adapter layer so the domain model never depends on the ORM.
type UserEntity struct {
	ID        string `gorm:"column:id;primaryKey"`
	FirstName string `gorm:"column:first_name;not null"`
	LastName  string `gorm:"column:last_name;not null"`
	Roles     string `gorm:"column:roles;not null;default:''"`
	Email     string `gorm:"column:email;not null;unique"`
	Password  string `gorm:"column:password;not null"`

	InstructorTitle string `gorm:"column:instructor_title;not null;default:''"`
	InstructorBio   string `gorm:"column:instructor_bio;not null;default:''"`

	Avatar string `gorm:"column:avatar;not null;default:''"`
}

func (UserEntity) TableName() string {
	return "users"
}

func fromDomain(u *user.User) *UserEntity {
	return &UserEntity{
		ID:              u.ID,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Roles:           rolesToString(u.Roles),
		Email:           u.Email,
		Password:        u.Password,
		InstructorTitle: u.InstructorTitle,
		InstructorBio:   u.InstructorBio,
		Avatar:          u.Avatar,
	}
}

func (m *UserEntity) toDomain() *user.User {
	return &user.User{
		ID:              m.ID,
		FirstName:       m.FirstName,
		LastName:        m.LastName,
		Roles:           rolesFromString(m.Roles),
		Email:           m.Email,
		Password:        m.Password,
		InstructorTitle: m.InstructorTitle,
		InstructorBio:   m.InstructorBio,
		Avatar:          m.Avatar,
	}
}

func rolesToString(roles []user.Role) string {
	parts := make([]string, len(roles))
	for i, role := range roles {
		parts[i] = string(role)
	}
	return strings.Join(parts, ",")
}

func rolesFromString(roles string) []user.Role {
	if roles == "" {
		return nil
	}

	parts := strings.Split(roles, ",")
	result := make([]user.Role, len(parts))
	for i, p := range parts {
		result[i] = user.Role(p)
	}
	return result
}
