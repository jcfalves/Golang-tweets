package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
	"time"
)

func (r *userRepository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*model.RefreshTokenModel, error) {
	// Implementation for retrieving a refresh token from the database
	query := "SELECT id, user_id, refresh_token, expired_at, created_at, updated_at FROM refresh_tokens WHERE user_id = ? AND expired_at > ?"
	row := r.db.QueryRowContext(ctx, query, userID, now)

	var refreshToken model.RefreshTokenModel
	err := row.Scan(&refreshToken.ID, &refreshToken.UserID, &refreshToken.RefreshToken, &refreshToken.ExpiresAt, &refreshToken.CreatedAt, &refreshToken.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No refresh token found for the user
		}
		return nil, err // Return error if query fails
	}

	return &refreshToken, nil // Return the retrieved refresh token
}
