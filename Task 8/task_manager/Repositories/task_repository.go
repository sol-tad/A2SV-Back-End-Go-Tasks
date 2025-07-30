package repository

import (
	"context"
	"time"
	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepo struct {
	collection *mongo.Collection
}

func NewTaskRepository(col *mongo.Collection) models.TaskRepository {
	return &taskRepo{collection: col}
}

func (r *taskRepo) GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *taskRepo) GetTaskByID(id string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}

	var task models.Task
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	return task, err
}

func (r *taskRepo) CreateTask(task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (r *taskRepo) UpdateTask(id string, updated models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return updated, err
	}

	update := bson.M{"$set": updated}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)

	updated.ID = objID
	return updated, err
}

func (r *taskRepo) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
