package user

import "context"

func (r *userRepository) DeleteRefreshTokenByUserID(ctx context.Context, userID int64) error {
	// Implementation for deleting refresh token by user ID
	query := "DELETE FROM refresh_tokens WHERE user_id = ?"
	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err // Return error if deletion fails
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err // Return error if unable to retrieve rows affected
	}

	if rowsAffected == 0 {
		return nil // No refresh token found for the user, nothing to delete
	}

	return nil
}
