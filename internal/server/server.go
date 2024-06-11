package server

import (
	"companypresence-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(userHandler *handlers.UserHandler) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,		
		AllowOriginsFunc: func(origin string) bool { return true },   		
		AllowHeaders: "Access-Control-Allow-Origin, Content-Type, Origin, Accept, Authorization",             
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		ExposeHeaders: "Content-Length",
	}))

	app.Use(logger.New())

	SetupRoutes(app, userHandler)

	return app
}

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler){
	app.Post("/login", userHandler.Login)

	api := app.Group("/api")	
	userRoutes := api.Group("/users")
	userRoutes.Post("/", userHandler.CreateUser)
	userRoutes.Get("/active", userHandler.GetActiveUsers)		
	userRoutes.Get("/email/:mail", userHandler.GetUserByMail)
	userRoutes.Get("/:id", userHandler.GetUserById)
	userRoutes.Get("/", userHandler.GetAllUsers)
	userRoutes.Put("/:id", userHandler.UpdateUser)
	userRoutes.Delete("/:id", userHandler.DeleteUser)
	
}