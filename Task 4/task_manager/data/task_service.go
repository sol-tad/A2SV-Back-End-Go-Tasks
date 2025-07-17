package data

import (
	"errors"
	"task_manager/models"
)

var tasks []models.Task
var nextID=1

func GetAllTasks() []models.Task{
	return tasks
}

func GetTaskByID(id int)(*models.Task,error){
	for _,t:=range tasks{
		if t.ID==id{
			return &t,nil
		}
	}
	return  nil,errors.New("task not found")
}

func CreateTask(task models.Task) models.Task{
	task.ID=nextID
	nextID++
	tasks=append(tasks, task)
	return task
}

func UpdateTask(id int,updated models.Task)(*models.Task,error){
	for i,t:=range tasks{
		if t.ID==id{
			updated.ID=id
			tasks[i]=updated
			return &updated,nil

		}
	}
	return nil,errors.New("task not found")
}

func DeleteTask(id int) error{
	for i,t:=range tasks{
		if t.ID==id{
			tasks=append(tasks[:i],tasks[i+1:]... )
			return nil
		}
	}
	return errors.New("task not found")
}