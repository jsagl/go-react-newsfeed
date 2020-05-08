package postgresql

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
	"github.com/lib/pq"
)


type PostgresUserStore struct {
	env *env.Env
	db *gorm.DB
}

func NewPostgresUserStore(env *env.Env, db *gorm.DB) storage.UserStore {
	return &PostgresUserStore{env: env, db: db}
}

func (store *PostgresUserStore) Create(user *models.User) (*models.User, error) {
	err := store.db.Create(&user).Error

	if err != nil {
		pqerr, ok := err.(*pq.Error)
		if ok && pqerr.Constraint == "uix_users_email" {
			e := models.NewErrBadParam(err, []string{"email"}, []string{"already_exists"})
			return nil, e
		}

		if ok && pqerr.Constraint == "uix_users_username" {
			e := models.NewErrBadParam(err, []string{"username"}, []string{"already_exists"})
			return nil, e
		}

		return nil, err
	}

	return user, nil
}

func (store *PostgresUserStore) FindByEmail(email string) (*models.User, error) {
	var user []*models.User

	err := store.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if len(user) == 0 {
		err := models.NewErrNotFound(errors.New("user_not_found"))
		return nil, err
	}

	return user[0], nil
}