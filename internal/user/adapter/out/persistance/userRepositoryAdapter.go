package persistance

import (
	"root_academy_go/internal/user/application/port/out"
	"root_academy_go/internal/user/domain/model"

	"gorm.io/gorm"
)

type UserRepositoryAdapter struct {
	db *gorm.DB
}

var _ out.IUserRepository = (*UserRepositoryAdapter)(nil)

func NewUserRepositoryAdapter(db *gorm.DB) *UserRepositoryAdapter {
	return &UserRepositoryAdapter{db: db}
}

func (r *UserRepositoryAdapter) SaveUser(user *model.User) error {
	entity := FromModel(user)
	return r.db.Create(&entity).Error
}

func (r *UserRepositoryAdapter) GetUserByEmail(email string) (*model.User, error) {
	var entity UserEntity
	if err := r.db.Where("email = ?", email).First(&entity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	user := ToModel(&entity)
	return &user, nil
}

func (r *UserRepositoryAdapter) GetUserByID(id string) (*model.User, error) {
	var entity UserEntity
	if err := r.db.Where("id = ?", id).First(&entity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	user := ToModel(&entity)
	return &user, nil
}

func (r *UserRepositoryAdapter) UpdateUser(id string) (*model.User, error) {
	var entity UserEntity
	if err := r.db.Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, err
	}
	if err := r.db.Save(&entity).Error; err != nil {
		return nil, err
	}
	user := ToModel(&entity)
	return &user, nil
}

func (r *UserRepositoryAdapter) DeleteUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&UserEntity{}).Error
}

func (r *UserRepositoryAdapter) GetUsers() []model.User {
	var entities []UserEntity
	r.db.Find(&entities)

	users := make([]model.User, 0, len(entities))
	for _, entity := range entities {
		users = append(users, ToModel(&entity))
	}
	return users
}
