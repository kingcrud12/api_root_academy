package persistance

import (
	"root_academy_go/internal/user/domain/model"
	"strings"
)

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

func FromModel(user *model.User) UserEntity {
	return UserEntity{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Roles:     RolesToString(user.Roles),
		Password:  user.Password,
	}

}

func ToModel(user *UserEntity) model.User {
	return model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Roles:     StringToRoles(user.Roles),
		Password:  user.Password,
	}

}

func RolesToString(roles []model.Role) string {
	strs := make([]string, len(roles))
	for i, r := range roles {
		strs[i] = string(r)
	}
	return strings.Join(strs, ",")
}

func StringToRoles(s string) []model.Role {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	roles := make([]model.Role, len(parts))
	for i, p := range parts {
		roles[i] = model.Role(p)
	}
	return roles
}
