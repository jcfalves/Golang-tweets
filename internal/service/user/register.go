package user

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error) {
	// Check if the user already exists by email or username
	existingUser, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	if existingUser != nil {
		return 0, http.StatusBadRequest, errors.New("user already exists")
	}

	//  hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	now := time.Now()

	// Create a new user model
	userModel := &model.UserModel{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(passwordHash),
		CreatedAt: now,
		UpdatedAt: now,
	}

	// create user
	userID, err := s.userRepo.CreateUser(ctx, userModel)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return userID, http.StatusCreated, nil

}
