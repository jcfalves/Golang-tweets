package main

import (
	"fmt"
	"go-tweets/internal/config"

	userHandler "go-tweets/internal/handler/user"
	userRepo "go-tweets/internal/repository/user"
	userService "go-tweets/internal/service/user"

	"go-tweets/pkg/internalsql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	// Validator
	validate := validator.New()

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := internalsql.ConnectMySql(cfg)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/check-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it1s working",
		})
	})

	userRepo := userRepo.NewRepository(db)
	userService := userService.NewService(cfg, userRepo)
	userHandler := userHandler.NewHandler(r, validate, userService)
	userHandler.RouteList(cfg.SecretJwt)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
