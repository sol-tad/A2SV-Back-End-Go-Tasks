package data

import (
	"context"
	"errors"
	"log"
	"task_manager/config"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Register(user models.User) error{
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//check if the username exists already
	var existing models.User
	err:=config.UserCollection.FindOne(ctx,bson.M{"username":user.Username}).Decode(&existing)
		if err == nil {
		return errors.New("username already exists")
	}
	if err != mongo.ErrNoDocuments {
		return err
	}
	//hash password
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil {
		return err
	}
	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "user"
	}

	_,err=config.UserCollection.InsertOne(ctx,user)

return err
}


func AuthenticateUser(username, password string) (models.User, error) {
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	var user models.User
	err := config.UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New("invalid credentials")
		}
		log.Println("Error finding user:", err)
		return models.User{}, err
	}

	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))

	if err!=nil {
		return models.User{}, errors.New("invalid credentials")
	}

	return user, nil
}