package main

import (
	"log"
	"net/http"

	"project/internal/config"
	"project/internal/db"
	"project/internal/handler"
	"project/internal/repository"
	"project/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Load config
	cfg := config.Load()

	// DB connection
	conn, err := db.NewPostgres(cfg.DBUrl)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// Dependency wiring
	userRepo := repository.NewUserRepository(conn)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	healthHandler := handler.NewHealthHandler()

	// Router
	r := chi.NewRouter()

	r.Get("/health", healthHandler.Health)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetUsers)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}