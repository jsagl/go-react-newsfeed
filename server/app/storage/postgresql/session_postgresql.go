package postgresql

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
	"time"
)


type PostgresSessionStore struct {
	env *env.Env
	db *gorm.DB
}

func NewPostgresSessionStore(env *env.Env, db *gorm.DB) storage.SessionStore {
	return &PostgresSessionStore{env: env, db: db}
}


func (store *PostgresSessionStore) StoreRememberMeToken(userId uint, token string) error {
	dbToken := &models.RememberMeToken{UserID: userId, Token: token, LastUsedAt: time.Now()}

	err := store.db.Create(&dbToken).Error

	if err != nil {
		return err
	}

	return nil
}

func (store *PostgresSessionStore) CheckRememberMeToken(token string) error {
	dbToken := &models.RememberMeToken{}
	store.db.Where("token = ?", token).First(&dbToken)
	if dbToken.ID == 0 {
		return errors.New("token_not_found")
	}

	err := store.db.Model(&dbToken).Update("last_used_at", time.Now()).Error
	if err != nil {
		return err
	}

	return nil
}

func (store *PostgresSessionStore) DestroyRememberMeToken(userId uint) error {
	err := store.db.Where("user_id = ?", userId).Delete(&models.RememberMeToken{}).Error

	if err != nil {
		return err
	}

	return nil
}