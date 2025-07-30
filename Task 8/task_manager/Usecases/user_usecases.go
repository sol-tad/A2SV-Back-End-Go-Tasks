package usecases

import (
	"errors"
	"task_manager/Domain"
	infrastructure "task_manager/Infrastructure"
	"task_manager/Repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(user models.User) error {
	_, err:=repository.FindUserByUsername(user.Username)
	if err==nil{
		return errors.New("username already exists")
	}
	if err!=mongo.ErrNoDocuments {
		return err
	}

	hashedPassword,err:=infrastructure.HashPassword(user.Password)
	if err!=nil{
		return err
	}
	user.Password=string(hashedPassword)

	if user.Role=="" {
		user.Role="user"
	}

	return repository.InsertUser(user)
}

func AuthenticateUser(username, password string) (models.User, error) {
	user,err:=repository.FindUserByUsername(username)
	if err!=nil {
		return models.User{}, errors.New("invalid credentials")
	}

	err=infrastructure.ComparePasswords(user.Password,password)
	if err!=nil {
		return models.User{}, errors.New("invalid credentials")
	}

	return user, nil
}
