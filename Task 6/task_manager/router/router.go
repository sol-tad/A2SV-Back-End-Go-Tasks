package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r:=gin.Default()
	r.POST("/login",controllers.Login)
	r.POST("/register",controllers.Register)
	r.GET("/tasks",controllers.GetTasks)

	auth:=r.Group("/",middleware.AuthMiddleware())
	{
		
		auth.POST("/tasks",controllers.CreateTask)
		// auth.GET("/tasks",controllers.GetTasks)
		auth.GET("/tasks/:id",controllers.GetTaskByID)
		auth.PUT("/task/:id",controllers.UpdateTask)
		auth.DELETE("/task/:id",middleware.AdminOnly(),controllers.DeleteTask)
	}


	return  r
}