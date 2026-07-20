package user

import (
	"go-tweets/internal/middleware"
	"go-tweets/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	// Define fields for the user handler, such as a user service

	api         *gin.Engine
	validate    *validator.Validate
	userService user.UserService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, userService user.UserService) *Handler {
	return &Handler{
		// Initialize fields as needed
		api:         api,
		validate:    validate,
		userService: userService,
	}
}

func (h *Handler) RouteList(secretkey string) {
	//h.api.GET("/users", h.ListUsers)

	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.Register)
	authRoute.POST("/login", h.Login)

	refreshTokenRoute := h.api.Group("/auth/refresh")
	refreshTokenRoute.Use(middleware.AuthRefreshTokenMiddleware(secretkey))
	refreshTokenRoute.POST("/refresh", h.RefreshToken)
}
