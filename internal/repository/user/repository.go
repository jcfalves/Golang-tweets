package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
	"time"
)

type UserRepository interface {
	// Define methods for user repository operations

	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, user *model.UserModel) (int64, error)
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*model.RefreshTokenModel, error)
	StoreRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error
	GetUserByID(ctx context.Context, userID int64) (*model.UserModel, error)
	DeleteRefreshTokenByUserID(ctx context.Context, userID int64) error
}

type userRepository struct {

	// Define fields for the user repository, such as a database connection
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		// Initialize fields as needed
		db: db,
	}
}
