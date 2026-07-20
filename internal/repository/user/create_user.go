package user

import (
	"context"
	"go-tweets/internal/model"
)

func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) (int64, error) {
	// Implementation for creating a new user in the database
	query := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err // Return error if insertion fails
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err // Return error if unable to retrieve last insert ID
	}

	return userID, nil // Return the newly created user's ID
}

/*
func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) (int64, error) {
	// Implementation for creating a new user in the database
	query := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err // Return error if insertion fails
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err // Return error if unable to retrieve last insert ID
	}

	return userID, nil // Return the newly created user's ID
}
*/
