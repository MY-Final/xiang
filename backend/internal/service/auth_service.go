package service

import (
	"errors"

	"xiang/backend/internal/config"
	"xiang/backend/internal/dto"
	jwtutil "xiang/backend/internal/pkg/jwt"
	"xiang/backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	adminRepo *repository.AdminRepository
	cfg       config.Config
}

func NewAuthService(adminRepo *repository.AdminRepository, cfg config.Config) *AuthService {
	return &AuthService{adminRepo: adminRepo, cfg: cfg}
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	admin, err := s.adminRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, expiresIn, err := jwtutil.Generate(s.cfg.JWTSecret, admin.ID, admin.Username, s.cfg.JWTExpireHours)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token:     token,
		ExpiresIn: expiresIn,
		User: dto.AdminProfile{
			ID:       admin.ID,
			Username: admin.Username,
			Nickname: admin.Nickname,
			Avatar:   admin.Avatar,
		},
	}, nil
}

func (s *AuthService) Me(adminID uint) (*dto.AdminProfile, error) {
	admin, err := s.adminRepo.FindByID(adminID)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, nil
	}

	return &dto.AdminProfile{
		ID:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Avatar:   admin.Avatar,
	}, nil
}
