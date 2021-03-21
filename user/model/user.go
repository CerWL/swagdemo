package model

import (
	"github.com/CerWL/swagdemo/db"
)

//Userlist is the respons of list user endpoint
type UserList struct {
	User []UserDEtail  `json:"user"`
}

//UserDetail detail model
type UserDetail struct{
	db.User
}

//CreateUser is create new user request model
type CreateUser struct {
	Name string `json:"name" minLength:"4"`
	Email string `json:"email" binding:"required"`
}

