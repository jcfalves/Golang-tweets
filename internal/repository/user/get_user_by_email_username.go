package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

func (r *userRepository) GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error) {
	// Implementation for getting user by email or username
	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = ? OR username = ?"
	row := r.db.QueryRowContext(ctx, query, email, username)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Other error occurred
	}

	return &user, nil
}
