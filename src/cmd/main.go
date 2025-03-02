package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/johancuervo/usersGO/config"
	"github.com/johancuervo/usersGO/internal/adapter/http"
	"github.com/johancuervo/usersGO/internal/adapter/repository"
	"github.com/johancuervo/usersGO/internal/adapter/routes"
	"github.com/johancuervo/usersGO/internal/service"
)

func main() {
	config.LoadEnv()
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal("❌ Error al conectar con la base de datos:", err)
	}
	defer db.Close()

	log.Println("🚀 Servidor iniciado...")
	userRepo := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)
	app := fiber.New()
	routes.RegisterRoutes(app, userHandler)

	log.Fatal(app.Listen(":8080"))
}
