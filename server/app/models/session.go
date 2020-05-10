package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type NewSessionForm struct {
	Email   	string 			`json:"email" binding:"required,email"`
	Password	string			`json:"password" binding:"required"`
	RememberMe	bool			`json:"RememberMe"`
}

type Session struct {
	UserID 		uint			`json:"user_id"`
	jwt.StandardClaims
}

type RememberMeToken struct {
	ID        	uint 			`gorm:"primary_key" json:"-"`
	Token   	string	 		`json:"token"`
	LastUsedAt 	time.Time		`json:"-"`
	UserID		uint			`gorm:"not null" json:"user_id" binding:"required"`
	User		User			`json:"-"`
	CreatedAt 	time.Time		`json:"created_at,omitempty"`
	UpdatedAt 	time.Time		`json:"updated_at,omitempty"`
	DeletedAt 	*time.Time		`json:"deleted_at,omitempty"`
}

func ConvertNewSessionFormToUser(input *NewSessionForm) *User {
	return &User{
		Email:    input.Email,
		Password: input.Password,
	}
}