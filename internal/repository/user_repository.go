package repository

import (
	"github.com/teakingwang/cursor-demo/internal/models"
	"github.com/teakingwang/cursor-demo/pkg/database"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create 创建用户
func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

// FindAll 获取所有用户
func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

// FindByID 根据ID获取用户
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

// Update 更新用户
func (r *UserRepository) Update(user *models.User) error {
	return database.DB.Save(user).Error
}

// Delete 删除用户
func (r *UserRepository) Delete(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}

// FindByUsername 根据用户名查找
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// FindByEmail 根据邮箱查找
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
