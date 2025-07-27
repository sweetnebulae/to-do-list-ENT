package router

import (
	"github.com/julienschmidt/httprouter"
	"os"
	"todo-list/controller"
	"todo-list/middleware"
	"todo-list/utils"
)

func NewRouter(authController *controller.UserController, taskController *controller.TaskController) *httprouter.Router {
	router := httprouter.New()
	cacheService := utils.NewCacheService()
	secretKey := os.Getenv("SECRET_KEY")

	authMiddleware := middleware.AuthMiddleware(secretKey, cacheService)

	// Public routes
	router.POST("/api/register", authController.Register)
	router.POST("/login", authController.Login)

	// Protected routes
	router.Handler("POST", "/logout", authMiddleware(middleware.Adapt(authController.Logout)))

	// Task Routes
	router.Handler("POST", "/tasks", authMiddleware(middleware.Adapt(taskController.Create)))
	router.Handler("GET", "/tasks", authMiddleware(middleware.Adapt(taskController.GetAll)))
	router.Handler("GET", "/tasks/:task_id", authMiddleware(middleware.Adapt(taskController.GetByID)))
	router.Handler("PUT", "/tasks/:task_id", authMiddleware(middleware.Adapt(taskController.Update)))
	router.Handler("DELETE", "/tasks/:task_id", authMiddleware(middleware.Adapt(taskController.Delete)))

	return router
}
