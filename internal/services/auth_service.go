package services

import (
	"errors"
	"instagram-clone/internal/models"
	"instagram-clone/internal/repository"
	"instagram-clone/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo     *repository.UserRepository
	jwtSecretKey string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecretKey string) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		jwtSecretKey: jwtSecretKey,
	}
}

func (s *AuthService) Register(username, email, password, fullName string) error {
	// Check if user exists
	if _, err := s.userRepo.FindByEmail(email); err == nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		FullName: fullName,
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, s.jwtSecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
