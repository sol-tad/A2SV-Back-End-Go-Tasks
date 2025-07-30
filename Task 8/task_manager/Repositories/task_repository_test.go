package repository_test

import (
	"testing"
	"task_manager/Domain"
	"task_manager/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepoTestSuite struct {
	suite.Suite
	mockRepo *mocks.TaskRepository
}

func (suite *TaskRepoTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.TaskRepository)
}

func (suite *TaskRepoTestSuite) TestCreateTask() {
	task := models.Task{
		Title:       "Test Task",
		Description: "Test Desc",
	}
	suite.mockRepo.On("CreateTask", task).Return(task, nil)

	result, err := suite.mockRepo.CreateTask(task)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepoTestSuite) TestGetTaskByID() {
	id := primitive.NewObjectID().Hex()
	expected := models.Task{Title: "Test", Description: "Test Desc"}
	suite.mockRepo.On("GetTaskByID", id).Return(expected, nil)

	result, err := suite.mockRepo.GetTaskByID(id)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepoTestSuite) TestDeleteTask() {
	id := primitive.NewObjectID().Hex()
	suite.mockRepo.On("DeleteTask", id).Return(nil)

	err := suite.mockRepo.DeleteTask(id)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}
