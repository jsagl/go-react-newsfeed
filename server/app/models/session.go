package models

import (
	"github.com/dgrijalva/jwt-go"
)

type NewSessionForm struct {
	Email   	string 			`json:"email" binding:"required,email"`
	Password	string			`json:"password" binding:"required"`
}

type Session struct {
	UserID 		uint			`json:"user_id"`
	jwt.StandardClaims
}

func ConvertNewSessionFormToUser(input *NewSessionForm) *User {
	return &User{
		Email:    input.Email,
		Password: input.Password,
	}
}