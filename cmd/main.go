package main

import (
	"hexagonal/internals/core/services"
	"hexagonal/internals/handlers"
	"hexagonal/internals/repositories"
	"hexagonal/internals/server"
)

func main() {
	mongoConn := "mongodb+srv://<been>:<yHGLdvjW4vp4HNzO>@ridem.pjvwk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	// repositories
	userRepository := repositories.NewUserRepository(mongoConn)
	// services
	userService := services.NewUserService(userRepository)
	// handlers
	userHandlers := handlers.NewUserHandlers(userService)
	// server
	httpServer := server.NewServer(
		userHandlers,
	)
	httpServer.Initialize()
}
