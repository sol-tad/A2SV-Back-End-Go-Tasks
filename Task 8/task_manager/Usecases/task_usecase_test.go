package usecases_test

import (
	"errors"
	"testing"
	"task_manager/Domain"
	"task_manager/mocks"
	"task_manager/Usecases"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

func TestGetAllTasks_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	mockTasks := []models.Task{
		{Title: "Task 1"},
		{Title: "Task 2"},
	}

	mockRepo.On("GetAllTasks").Return(mockTasks, nil)

	usecase := usecases.NewTaskUseCase(mockRepo)

	tasks, err := usecase.GetAllTasks()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(tasks))
	assert.Equal(t, "Task 1", tasks[0].Title)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	mockRepo.On("GetTaskByID", "invalid-id").Return(models.Task{}, errors.New("not found"))

	usecase := usecases.NewTaskUseCase(mockRepo)

	_, err := usecase.GetTaskByID("invalid-id")

	assert.Error(t, err)
	assert.Equal(t, "task not found", err.Error())
	mockRepo.AssertExpectations(t)
}
