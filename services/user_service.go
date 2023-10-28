package services

import (
	"code-ai/models/domain"
	"code-ai/repository"
	"fmt"
)

type UserService struct {
	UserRepository repository.UserRepository
}

type UserServiceImp interface {
	CreateUser(user *domain.User) (*domain.User, error)
	FindUserByID(id int64) (*domain.User, error)
	ReduceLimitText(userID int64, userTextCount int64) (*domain.User, error)
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (use *UserService) FindUserByID(id int64) (*domain.User, error) {
	user, err := use.UserRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (use *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	fmt.Println("test")
	user, err := use.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (use *UserService) ReduceLimitText(userID int64, userTextCount int64) (*domain.User, error) {
	user, err := use.UserRepository.ReduceLimitText(userID, userTextCount)
	if err != nil {
		return nil, err
	}
	return user, nil
}
