package router

import (
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
	"task_manager/Usecases"
	"task_manager/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	// Set up dependencies
	taskRepo := repository.NewTaskRepository(config.TaskCollection)
	taskUsecase := usecases.NewTaskUseCase(taskRepo)
	taskController := controllers.NewTaskController(taskUsecase)

	// Protected routes with middleware
	auth := r.Group("/", infrastructure.AuthMiddleware())
	{
		auth.POST("/tasks", taskController.CreateTask)
		auth.GET("/tasks", taskController.GetTasks)
		auth.GET("/tasks/:id", taskController.GetTaskByID)
		auth.PUT("/tasks/:id", taskController.UpdateTask)
		auth.DELETE("/tasks/:id", infrastructure.AdminOnly(), taskController.DeleteTask)
	}

	return r
}
