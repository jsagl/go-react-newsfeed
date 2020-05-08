package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
	"golang.org/x/crypto/bcrypt"
)

type SessionUsecaseInterface interface {
	Create(c *gin.Context) (*models.User, error)
}

type SessionUsecase struct {
	env *env.Env
	store storage.UserStore
}

func NewSessionUsecase(env *env.Env, db storage.UserStore) SessionUsecaseInterface {
	return &SessionUsecase{env: env, store: db}
}

func (usecase *SessionUsecase) Create(c *gin.Context) (*models.User, error) {
	var loginForm models.NewSessionForm

	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		// Will be necessary to provide more information to user on input error
		e := models.NewErrBadParam(err, []string{}, []string{})
		return nil, e
	}

	user, err := usecase.store.FindByEmail(loginForm.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginForm.Password)); err != nil {
		e := models.NewErrBadParam(err, []string{"password"}, []string{"invalid_password"})
		return nil, e
	}

	return user, nil
}