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
	userService := service.NewUserService(client, secretKey, cacheService)
	userController := controller.NewUserController(*userService, *cacheService)
	routes := router.NewRouter(userController)

	config.StartServer(routes)
	select {}
}
