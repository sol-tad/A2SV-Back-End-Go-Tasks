package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, data.GetAllTasks())
}

func GetTaskByID(c *gin.Context){

		id,err:=strconv.Atoi(c.Param("id"))

		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"invalid task id"})
			return
		}
		task,err:=data.GetTaskByID(id)
		if err!=nil{
			c.JSON(http.StatusNotFound,gin.H{"error":"task not found"})
			return
		}
		c.JSON(http.StatusOK,task)

}		

func CreateTask(c *gin.Context){
	var task models.Task
	if err:=c.ShouldBindJSON(&task);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	newTask:=data.CreateTask(task)
	c.JSON(http.StatusCreated,newTask)

}

func UpdateTask(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var task models.Task
	if err:=c.ShouldBindJSON(&task);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	updatedTask,err:=data.UpdateTask(id,task)

		if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK,updatedTask)

}

func DeleteTask(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err:= data.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, "deleted successfully")

}