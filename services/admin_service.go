package services

import (
	"code-ai/models/domain"
	"code-ai/models/web"
	"code-ai/repository"
	"code-ai/utils/helper"
	"code-ai/utils/req"
	"fmt"

	"github.com/go-playground/validator"
)

type AdminService struct {
	AdminRepo repository.AdminRepository
	Validator *validator.Validate
}

type AdminServiceInterface interface {
	CreateAdminService(request web.AdminCreateRequest) (*domain.Admin, error)
	LoginAdminService(request web.AdminLoginRequest) (*domain.Admin, error)
	FindAllAdminService() ([]domain.Admin, error)
	FindAdminByEmailService(email string) (*domain.Admin, error)
	FindAllUserService() ([]domain.User, error)
	FindUserByUsernameService(username string) (*domain.User, error)
	UpdateUserService(userID int, user web.UserUpdateRequest) (*domain.User, error)
	DeleteUserService(userID int) error
}

func NewAdminService(repo repository.AdminRepository, validate *validator.Validate) *AdminService {
	return &AdminService{
		AdminRepo: repo,
		Validator: validate,
	}
}

func (as *AdminService) CreateAdminService(request web.AdminCreateRequest) (*domain.Admin, error) {
	fmt.Println("line________________________________________")
	fmt.Println("line2________________________________________")
	existingAdmin, _ := as.AdminRepo.FindAdminByEmail(request.Email)
	if existingAdmin != nil {
		return nil, fmt.Errorf("email already exist")
	}

	admin := req.AdminCreateRequestToAdminDomain(request)

	admin.Password = helper.HashPassword(admin.Password)

	result, err := as.AdminRepo.CreateAdmin(admin)
	if err != nil {
		return nil, fmt.Errorf("error when creating Admin: %s", err.Error())
	}

	return result, nil
}

func (as *AdminService) LoginAdminService(request web.AdminLoginRequest) (*domain.Admin, error) {
	
	existingAdmin, err := as.AdminRepo.FindAdminByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email or Password")
	}

	admin := req.AdminLoginRequestToAdminDomain(request)

	err = helper.ComparePassword(existingAdmin.Password, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingAdmin, nil
}

func (as *AdminService) FindAllAdminService() ([]domain.Admin, error) {
	admin, err := as.AdminRepo.FindAllAdmin()
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}
	return admin, nil
}


func (as *AdminService) FindAdminByEmailService(email string) (*domain.Admin, error) {
	admin, err := as.AdminRepo.FindAdminByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}
	return admin, nil
}

func (as *AdminService) FindUserByUsernameService(username string) (*domain.User, error) {
	user, err := as.AdminRepo.FindUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}
	return user, nil
}

func (as *AdminService) FindAllUserService() ([]domain.User, error) {
	user, err := as.AdminRepo.FindAllUser()
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}
	return user, nil
}

func (as *AdminService) UpdateUserService(userID int, user web.UserUpdateRequest) (*domain.User, error) {
	userUpdate := req.UserUpdateRequestToUserDomain(user)
	result, err := as.AdminRepo.UpdateUser(userID, userUpdate)
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}
	return result, nil
}

func (as *AdminService) DeleteUserService(userID int) error {
	err := as.AdminRepo.DeleteUser(userID)
	if err != nil {
		return fmt.Errorf("admins not found")
	}
	return nil
}
