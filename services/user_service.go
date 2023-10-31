package services

import (
	"code-ai/models/domain"
	"code-ai/repository"
	"fmt"
)

type UserService struct {
	UserRepository repository.UserRepository
}

type UserServiceInterface interface {
	CreateUser(user *domain.User) (*domain.User, error)
	FindUserByID(id int64) (*domain.User, error)
	ReduceLimitText(userID int64, userTextCount int64) (*domain.User, error)
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (us *UserService) FindUserByID(id int64) (*domain.User, error) {
	user, err := us.UserRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	fmt.Println("test")
	user, err := us.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) ReduceLimitText(userID int64, userCredit int64) (*domain.User, error) {
	user, err := us.UserRepository.ReduceLimitText(userID, userCredit)
	if err != nil {
		return nil, err
	}
	return user, nil
}
