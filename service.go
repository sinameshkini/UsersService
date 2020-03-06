package main

import (
	"crudservice/models"
	"fmt"
)

type UserService interface {
	Fetch() ([]*models.User, error)
	Create(user *models.User) (*models.User, error)
}

type userService struct {}

var users []*models.User

func (userService) Fetch() ([]*models.User, error){
	return users, nil
}

func (userService) Create(user *models.User) (*models.User, error){
	fmt.Println(user)
	users = append(users, user)
	return user, nil
}

//var ErrEmpty =