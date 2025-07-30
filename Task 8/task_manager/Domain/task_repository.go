package models

type TaskRepository interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	DeleteTask(id string) error
}
