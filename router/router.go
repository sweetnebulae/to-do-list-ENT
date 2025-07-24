package router

import (
	"github.com/julienschmidt/httprouter"
	"os"
	"todo-list/controller"
	"todo-list/middleware"
	"todo-list/utils"
)

func NewRouter(authController *controller.UserController) *httprouter.Router {
	router := httprouter.New()
	cacheService := utils.NewCacheService()
	secretKey := os.Getenv("SECRET_KEY")

	authMiddleware := middleware.AuthMiddleware(secretKey, cacheService)

	// Public routes
	router.POST("/api/register", authController.Register)
	router.POST("/login", authController.Login)

	// Protected routes (gunakan middleware lalu bridge ke httprouter)
	router.Handler("POST", "/logout", authMiddleware(middleware.Adapt(authController.Logout)))

	return router
}
