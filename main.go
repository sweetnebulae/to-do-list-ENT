package main

import (
	"os"
	"todo-list/config"
	"todo-list/controller"
	"todo-list/router"
	"todo-list/service"
	"todo-list/utils"
)

func main() {
	client := config.ConnectDB()
	defer config.DisconnectDB(client)

	secretKey := os.Getenv("SECRET_KEY")
	cacheService := utils.NewCacheService()

	// Initialising User Service and Controller
	userService := service.NewUserService(client, secretKey, cacheService)
	userController := controller.NewUserController(*userService, *cacheService)

	// Initialising Task Service and Controller
	taskService := service.NewTaskService(client)
	taskController := controller.NewTaskController(*taskService)

	// Inject to all route
	routes := router.NewRouter(userController, taskController)

	// Start server
	config.StartServer(routes)

	select {}
}
