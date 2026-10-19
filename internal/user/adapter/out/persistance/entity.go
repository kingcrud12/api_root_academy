package persistance

import "root_academy_go/internal/user/domain/model"

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
		Password:  user.Password,
	}

}

func ToModel(user *UserEntity) model.User {
	return model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

}
