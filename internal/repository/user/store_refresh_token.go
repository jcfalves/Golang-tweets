package user

import (
	"context"
	"go-tweets/internal/model"
)

func (r *userRepository) StoreRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error {
	query := "INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiresAt, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return err // Return error if insertion fails
	}

	return nil // Return nil if insertion is successful
}
