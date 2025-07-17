package router

import (
	"task_manager/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r:=gin.Default()
	r.GET("/tasks",controllers.GetTasks)
	r.GET("/tasks/:id",controllers.GetTaskByID)
	r.POST("tasks",controllers.CreateTask)
	r.PUT("/task/:id",controllers.UpdateTask)
	r.DELETE("/task/:id",controllers.DeleteTask)


	return  r
}