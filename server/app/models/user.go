package models

import "time"

type User struct {
	ID        	uint 			`gorm:"primary_key" json:"id,omitempty"`
	Email   	string 			`gorm:"type:varchar(100);unique_index;not null" json:"email"`
	Username 	string			`gorm:"type:varchar(100);unique_index;not null" json:"username"`
	Password	string			`gorm:"type:varchar(100);not null" json:"-"`
	Favorites   []Favorite		`json:"-"`
	CreatedAt 	time.Time		`json:"created_at,omitempty"`
	UpdatedAt 	time.Time		`json:"updated_at,omitempty"`
	DeletedAt 	*time.Time		`json:"deleted_at,omitempty"`
}

type NewUserForm struct {
	Email   	string 			`json:"Email" binding:"required,email"`
	Username 	string			`json:"Username" binding:"required,min=3,max=50"`
	Password	string			`json:"Password" binding:"required,min=8,max=45"`
}

func ConvertNewUserFormToUser(input *NewUserForm) *User {
	return &User{
		Email: input.Email,
		Username: input.Username,
		Password: input.Password,
	}
}