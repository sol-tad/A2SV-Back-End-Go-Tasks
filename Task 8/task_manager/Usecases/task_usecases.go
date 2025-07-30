package usecases

import (
	"errors"
	"task_manager/Domain"
)

type TaskUseCase struct {
	Repo models.TaskRepository
}

// NewTaskUseCase creates a new TaskUseCase with injected repository
func NewTaskUseCase(repo models.TaskRepository) *TaskUseCase {
	return &TaskUseCase{Repo: repo}
}

func (uc *TaskUseCase) GetAllTasks() ([]models.Task, error) {
	return uc.Repo.GetAllTasks()
}

func (uc *TaskUseCase) GetTaskByID(id string) (models.Task, error) {
	task, err := uc.Repo.GetTaskByID(id)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (uc *TaskUseCase) CreateTask(task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("title is required")
	}
	return uc.Repo.CreateTask(task)
}

func (uc *TaskUseCase) UpdateTask(id string, updated models.Task) (models.Task, error) {
	if updated.Title == "" {
		return models.Task{}, errors.New("title is required")
	}
	return uc.Repo.UpdateTask(id, updated)
}

func (uc *TaskUseCase) DeleteTask(id string) error {
	return uc.Repo.DeleteTask(id)
}
