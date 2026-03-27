package main

import (
	"github.com/SaviolaX/blog/internal/config"
	"github.com/SaviolaX/blog/internal/db"
	"github.com/SaviolaX/blog/internal/handler"
	"github.com/SaviolaX/blog/internal/repository"
	"github.com/SaviolaX/blog/internal/service"
)

func main() {
	cfg := config.Load()

	database := db.Connect(cfg.DB)
	db.Migrate(database)

	userRepo := repository.NewUserRepository(database)
	postRepo := repository.NewPostRepository(database)

	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)

	userHandler := handler.NewUserHandler(cfg.JWT.ExpireHours, cfg.JWT.Secret, userService)
	postHandler := handler.NewPostHandler(postService)

	r := handler.SetupRouter(userHandler, postHandler, cfg.JWT.Secret)

	r.Run(cfg.App.Port)
}
