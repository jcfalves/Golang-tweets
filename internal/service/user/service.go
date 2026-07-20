package user

import (
	"context"
	"go-tweets/internal/config"
	"go-tweets/internal/dto"
	"go-tweets/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)
	RefreshToken(ctx context.Context, req *dto.RefresehTokenRequest, userID int64) (string, string, int, error)
}

type userService struct {
	// Define fields for the user service, such as a user repository
	cfg      *config.Config
	userRepo user.UserRepository
}

// NewUserService creates a new instance of UserService
func NewService(cfg *config.Config, userRepo user.UserRepository) UserService {
	return &userService{
		// Initialize fields as needed
		cfg:      cfg,
		userRepo: userRepo,
	}
}
