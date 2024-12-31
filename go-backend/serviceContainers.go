package main

import (
	"go-backend/controllers"
	"go-backend/interfaces"
	"go-backend/repositories"
	"go-backend/services"
)

func InjectUsersController() interfaces.IUserController {
	usersRepo := &repositories.UserRepository{}
	usersService := &services.UserService{IUserRepository: usersRepo}
	usersController := &controllers.UserController{IUserService: usersService}

	return usersController
}
