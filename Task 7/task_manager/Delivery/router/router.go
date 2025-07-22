package router

import (
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r:=gin.Default()
	r.POST("/login",controllers.Login)
	r.POST("/register",controllers.Register)

	auth:=r.Group("/",infrastructure.AuthMiddleware())
	{
		auth.POST("/tasks",controllers.CreateTask)
		auth.GET("/tasks",controllers.GetTasks)
		auth.GET("/tasks/:id",controllers.GetTaskByID)
		auth.PUT("/task/:id",controllers.UpdateTask)
		auth.DELETE("/task/:id",infrastructure.AdminOnly(),controllers.DeleteTask)
	}


	return  r
}