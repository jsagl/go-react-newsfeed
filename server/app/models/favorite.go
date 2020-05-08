package models

import (
	"time"
)

type Favorite struct {
	ID        	uint 			`gorm:"primary_key" json:"id,omitempty" binding:"-"`
	Title   	*string 		`json:"title" binding:"required"`
	TargetUrl 	*string			`json:"target_url"`
	Date		*time.Time		`json:"date"`
	Category	*string			`json:"category"`
	Tags		*string			`json:"tags"`
	SourceName	*string			`json:"source_name"`
	SourceUrl	*string			`json:"source_url"`
	UserID		uint			`gorm:"not null" json:"user_id" binding:"required"`
	User		User			`json:"-"`
	CreatedAt 	time.Time		`json:"created_at,omitempty"`
	UpdatedAt 	time.Time		`json:"updated_at,omitempty"`
	DeletedAt 	*time.Time		`json:"deleted_at,omitempty"`
}

type FavoriteInput struct {
	Title   	*string 		`json:"title" binding:"required"`
	TargetUrl 	*string			`json:"target_url" binding:"required"`
	Date		*time.Time		`json:"date"`
	Category	*string			`json:"category"`
	Tags		*string			`json:"tags"`
	SourceUrl	*string			`json:"source_url"`
    SourceName	*string			`json:"source_name"`
}

type FavoriteDestroy struct {
	TargetUrl 	*string			`json:"target_url" binding:"required"`
}

func ConvertToFavorite(input *FavoriteInput) *Favorite {
	return &Favorite{
		Title: input.Title,
		TargetUrl: input.TargetUrl,
		Date: input.Date,
		Category: input.Category,
		Tags: input.Tags,
		SourceName: input.SourceName,
		SourceUrl: input.SourceUrl,
	}
}