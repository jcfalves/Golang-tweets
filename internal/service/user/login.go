package user

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"go-tweets/pkg/jwt"
	"go-tweets/pkg/refreshtoken"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {

	// Get user by email or username
	user, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	if user == nil {
		return "", "", http.StatusNotFound, errors.New("invalid credentials")
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", "", http.StatusUnauthorized, errors.New("invalid credentials")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	now := time.Now()
	refreshToken, err := s.userRepo.GetRefreshToken(ctx, user.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshToken != nil {
		return token, refreshToken.RefreshToken, http.StatusOK, nil
	}

	refreshTokenString, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshTokenString,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour), // Set expiration time for the refresh token (e.g., 7 days)
		CreatedAt:    now,
		UpdatedAt:    now,
	})

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return token, refreshTokenString, http.StatusOK, nil

	// Generate access and refresh tokens (implement your token generation logic here)

}
