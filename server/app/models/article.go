package models

import (
	"time"
)

type Article struct {
	Title   	string		`json:"title,omitempty"`
	TargetUrl 	string		`json:"target_url,omitempty"`
	Date		time.Time	`json:"date,omitempty"`
	Category	string		`json:"category,omitempty"`
	SourceUrl	string		`json:"source_url,omitempty"`
	SourceName	string		`json:"source_name,omitempty"`
	Bookmarked	bool		`json:"bookmarked,omitempty"`
}