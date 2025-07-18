package data

import (
	"context"
	"log"
	"task_manager/config"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Decode error:", err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	var task models.Task
	err := config.TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)

	return task, err
}

func CreateTask(task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := config.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID)
	return task, nil
}

func UpdateTask(id string, updated models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{
		"$set": updated,
	}
	_, err := config.TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	updated.ID = objID
	return updated, err
}

func DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := config.TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
