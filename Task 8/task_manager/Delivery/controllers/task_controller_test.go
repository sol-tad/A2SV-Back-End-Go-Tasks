package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"task_manager/Domain"
	"task_manager/Delivery/controllers"
	"task_manager/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(controller *controllers.TaskController) *gin.Engine {
	r := gin.Default()
	r.GET("/tasks", controller.GetTasks)
	r.GET("/tasks/:id", controller.GetTaskByID)
	r.POST("/tasks", controller.CreateTask)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)
	return r
}

func TestTaskController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.TaskUsecase)
	controller := controllers.NewTaskController(mockUsecase)
	r := setupRouter(controller)

	exampleTask := models.Task{ID: [12]byte{}, Title: "Sample Task", Description: "Test"}
	tasks := []models.Task{exampleTask}

	mockUsecase.On("GetAllTasks").Return(tasks, nil)
	t.Run("GetTasks", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/tasks", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	mockUsecase.On("GetTaskByID", "123").Return(exampleTask, nil)
	t.Run("GetTaskByID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/tasks/123", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	mockUsecase.On("CreateTask", mock.AnythingOfType("models.Task")).Return(exampleTask, nil)
	t.Run("CreateTask", func(t *testing.T) {
		payload, _ := json.Marshal(exampleTask)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		assert.Equal(t, 201, w.Code)
	})

	mockUsecase.On("UpdateTask", "123", mock.AnythingOfType("models.Task")).Return(exampleTask, nil)
	t.Run("UpdateTask", func(t *testing.T) {
		payload, _ := json.Marshal(exampleTask)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/tasks/123", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	mockUsecase.On("DeleteTask", "123").Return(nil)
	t.Run("DeleteTask", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/tasks/123", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	mockUsecase.AssertExpectations(t)
}
