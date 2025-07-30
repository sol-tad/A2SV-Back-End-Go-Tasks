package repository

import (
	"context"
	"task_manager/config"
	"task_manager/Domain"
	"time"
	"go.mongodb.org/mongo-driver/bson"
)

func FindUserByUsername(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := config.UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}

func InsertUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.UserCollection.InsertOne(ctx, user)
	return err
}
