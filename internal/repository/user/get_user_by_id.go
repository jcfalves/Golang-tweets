package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*model.UserModel, error) {
	// Implementation for retrieving a user by ID from the database
	query := "SELECT id, email, username, password, created_at, updated_at FROM users WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, userID)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found with the given ID
		}
		return nil, err // Return error if query fails
	}

	return &user, nil // Return the retrieved user
}
