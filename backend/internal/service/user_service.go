package service

import (
	"context"
	"errors"
	"time"
	"xiang/backend/internal/config"
	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"
	jwtutil "xiang/backend/internal/pkg/jwt"
	"xiang/backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists     = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUserInvalidCredential = errors.New("invalid username or password")
)

type UserService struct {
	userRepo *repository.UserRepository
	cfg      config.Config
}

func NewUserService(userRepo *repository.UserRepository, cfg config.Config) *UserService {
	return &UserService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

// Register 用户注册
func (s *UserService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserProfile, error) {
	// 检查用户名是否存在
	exists, err := s.userRepo.UsernameExists(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	// 检查邮箱是否存在
	if req.Email != "" {
		exists, err = s.userRepo.EmailExists(ctx, req.Email)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrEmailAlreadyExists
		}
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 解析日期
	var birthday, anniversary *time.Time
	if req.Birthday != "" {
		if t, err := time.Parse("2006-01-02", req.Birthday); err == nil {
			birthday = &t
		}
	}
	if req.Anniversary != "" {
		if t, err := time.Parse("2006-01-02", req.Anniversary); err == nil {
			anniversary = &t
		}
	}

	// 创建用户
	user := &model.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Nickname:     req.Nickname,
		Email:        req.Email,
		Birthday:     birthday,
		Anniversary:  anniversary,
		Status:       "active",
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserProfile{
		ID:          user.ID,
		Username:    user.Username,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		Email:       user.Email,
		Birthday:    req.Birthday,
		Anniversary: req.Anniversary,
	}, nil
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context, req dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	user, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserInvalidCredential
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrUserInvalidCredential
	}

	// 生成 Token
	token, expiresIn, err := jwtutil.Generate(s.cfg.JWTSecret, user.ID, user.Username, s.cfg.JWTExpireHours)
	if err != nil {
		return nil, err
	}

	return &dto.UserLoginResponse{
		Token:     token,
		ExpiresIn: expiresIn,
		User: dto.UserProfile{
			ID:          user.ID,
			Username:    user.Username,
			Nickname:    user.Nickname,
			Avatar:      user.Avatar,
			Email:       user.Email,
			Birthday:    user.Birthday.Format("2006-01-02"),
			Anniversary: user.Anniversary.Format("2006-01-02"),
		},
	}, nil
}

// GetProfile 获取用户资料
func (s *UserService) GetProfile(ctx context.Context, userID uint) (*dto.UserProfile, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	profile := &dto.UserProfile{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
	}

	if user.Birthday != nil {
		profile.Birthday = user.Birthday.Format("2006-01-02")
	}
	if user.Anniversary != nil {
		profile.Anniversary = user.Anniversary.Format("2006-01-02")
	}

	return profile, nil
}

// UpdateProfile 更新用户资料
func (s *UserService) UpdateProfile(ctx context.Context, userID uint, req dto.RegisterRequest) (*dto.UserProfile, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 更新资料
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Email != "" && req.Email != user.Email {
		// 检查新邮箱是否已被使用
		exists, err := s.userRepo.EmailExists(ctx, req.Email)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrEmailAlreadyExists
		}
		user.Email = req.Email
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return s.GetProfile(ctx, userID)
}
