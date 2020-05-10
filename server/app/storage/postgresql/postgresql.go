package postgresql

import (
	"github.com/jinzhu/gorm"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"os"
)

func NewPostgresDatabase() (*gorm.DB, error) {
	// Parse postgres db url
	args, _ := pq.ParseURL(os.Getenv("DATABASE_URL"))
	dbParams := args + " connect_timeout=" + os.Getenv("TIMEOUT_IN_SEC")

	// Open postgresql database
	db, err := gorm.Open("postgres", dbParams)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Favorite{}, &models.User{}, &models.RememberMeToken{})

	return db, nil
}