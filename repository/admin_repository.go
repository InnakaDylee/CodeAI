package repository

import (
	"code-ai/models/domain"
	"code-ai/utils/req"
	"code-ai/utils/res"
	"fmt"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

type AdminRepositoryInt interface {
	CreateAdmin(admin *domain.Admin) (*domain.Admin, error)
	FindAllAdmin() (*domain.Admin, error)
	FindAdminByEmail(email string) (*domain.Admin, error)
	FindUserByID(userID int64) (*domain.User, error)
	FindAllUser() ([]*domain.User, error)
	UpdateUser(userID int64, user *domain.User) (*domain.User, error)
	DeleteUser(userID int64) error
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) CreateAdmin(admin *domain.Admin) (*domain.Admin, error) {
	adminRepo := req.AdminDomaintoAdminSchema(*admin)
	result := r.db.Create(&adminRepo)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.AdminSchemaToAdminDomain(adminRepo)

	return results, nil
}

func (r *AdminRepository) FindAllAdmin() ([]domain.Admin, error) {
	admin := []domain.Admin{}
	err := r.db.Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepository) FindAdminByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	fmt.Println(email)
	err := r.db.First(&admin, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) FindUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Message").First(&user, "name = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AdminRepository) FindAllUser() ([]domain.User, error) {
	var user []domain.User
	err := r.db.Preload("Message").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AdminRepository) UpdateUser(userID int, user *domain.User) (*domain.User, error) {
	err := r.db.Model(&user).Where("id = ?", userID).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AdminRepository) DeleteUser(userID int) error {
	var user domain.User
	fmt.Println(userID)
	err := r.db.Delete(&user, "id = ?", userID).Error
	if err != nil {
		return  err
	}
	return nil
}