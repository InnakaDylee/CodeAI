package repository

import (
	"code-ai/models/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

type UserRepositoryImp interface {
	FindUserByID(userID int64) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	AddLimitText(userID int64) (*domain.User, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{
        db,
    }
}

func (r *UserRepository) FindUserByID(userID int64) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) AddLimitText(userID int64,userTextCount int64) (*domain.User, error) {
	var user domain.User
	err := r.db.Model(&user).Where("id = ?", userID).Update("text_count", userTextCount + 1).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}