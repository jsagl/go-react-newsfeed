package postgresql

import (
	"github.com/jinzhu/gorm"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
	"time"
)


type PostgresFavoriteStore struct {
	env *env.Env
	db *gorm.DB
}

func NewPostgresFavoriteStore(env *env.Env, db *gorm.DB) storage.FavoriteStore {
	return &PostgresFavoriteStore{env: env, db: db}
}

func (store *PostgresFavoriteStore) Index(userId uint) ([]*models.Favorite, error) {
	var favorites []*models.Favorite
	err := store.db.Order("date desc").Where("user_id = ?", userId).Find(&favorites).Error

	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func (store *PostgresFavoriteStore) Create(favorite *models.Favorite) (*models.Favorite, error) {
	err := store.db.Create(&favorite).Error

	if err != nil {
		return nil, err
	}

	return favorite, nil
}

func (store *PostgresFavoriteStore) DestroyByTargetUrlAndUserId(targetUrl string, userId uint) error {
	query := "DELETE FROM favorites WHERE target_url = $1 AND user_id = $2"
	err := store.db.Exec(query, targetUrl, userId).Error

	if err != nil {
		return err
	}

	return nil
}

func (store *PostgresFavoriteStore) BookmarkedUrls(userId uint) ([]string, error) {
	var favorites []*models.Favorite
	var urls []string

	date := time.Now().AddDate(0, -1, 0)
	err := store.db.Where("created_at > ? AND user_id = ?", date, userId).Find(&favorites).Pluck("target_url", &urls).Error

	if err != nil {
		return nil, err
	}

	return urls, nil
}