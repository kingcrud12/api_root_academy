package postgres

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"root_academy_go/internal/domain/model/user"
	"root_academy_go/internal/domain/repository"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) SaveUser(u *user.User) error {
	entity := fromDomain(u)

	if err := r.db.Create(entity).Error; err != nil {
		return fmt.Errorf("create user %s: %w", u.ID, err)
	}

	return nil
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*user.User, error) {
	var entity UserEntity

	err := r.db.Where("email = ?", email).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get user by email %s: %w", email, err)
	}

	return entity.toDomain(), nil
}

func (r *UserRepositoryImpl) GetUserByID(id string) (*user.User, error) {
	var entity UserEntity

	err := r.db.Where("id = ?", id).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get user by id %s: %w", id, err)
	}

	return entity.toDomain(), nil
}

func (r *UserRepositoryImpl) UpdateUser(id string) (*user.User, error) {
	var entity UserEntity

	if err := r.db.Where("id = ?", id).First(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("update user %s: %w", id, err)
	}

	return entity.toDomain(), nil
}

func (r *UserRepositoryImpl) DeleteUser(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&UserEntity{}).Error; err != nil {
		return fmt.Errorf("delete user %s: %w", id, err)
	}

	return nil
}

func (r *UserRepositoryImpl) GetUsers() []user.User {
	var entities []UserEntity

	if err := r.db.Find(&entities).Error; err != nil {
		return nil
	}

	users := make([]user.User, len(entities))
	for i, e := range entities {
		users[i] = *e.toDomain()
	}

	return users
}
