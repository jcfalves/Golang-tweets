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
)

func (s *userService) RefreshToken(ctx context.Context, req *dto.RefresehTokenRequest, userID int64) (string, string, int, error) {

	// Retrieve the refresh token from the repository
	userExist, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExist == nil {
		return "", "", http.StatusNotFound, errors.New("user not found") // User not found
	}

	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExists == nil {
		return "", "", http.StatusNotFound, errors.New("refresh token was expired") // Refresh token not found
	}

	if refreshTokenExists.RefreshToken != req.RefreshToken {
		return "", "", http.StatusUnauthorized, errors.New("invalid refresh token OR not found") // Invalid refresh token
	}

	token, err := jwt.CreateToken(userID, userExist.Username, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepo.DeleteRefreshTokenByUserID(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	now := time.Now()
	s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour), // Set expiration time for the refresh token (e.g., 7 days)
		CreatedAt:    now,
		UpdatedAt:    now,
	})

	return token, refreshToken, http.StatusOK, nil // Return success response (you can modify this as needed)
}
