package usecases

import (
	"errors"
	"task_manager/Domain"
	"task_manager/Repositories"
)

func GetAllTasks() ([]models.Task, error) {
	return repository.GetAllTasks()
}

func GetTaskByID(id string) (models.Task, error) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func CreateTask(task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("title is required")
	}
	return repository.CreateTask(task)
}

func UpdateTask(id string, updated models.Task) (models.Task, error) {
	if updated.Title == "" {
		return models.Task{}, errors.New("title is required")
	}
	return repository.UpdateTask(id, updated)
}

func DeleteTask(id string) error {
	return repository.DeleteTask(id)
}
