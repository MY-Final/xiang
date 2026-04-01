package repository

import (
	"context"
	"xiang/backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

func (r *UserRepository) UsernameExists(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
