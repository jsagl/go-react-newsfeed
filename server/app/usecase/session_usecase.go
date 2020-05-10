package usecase

import (
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
	"golang.org/x/crypto/bcrypt"
)

type SessionUsecaseInterface interface {
	Create(form models.NewSessionForm) (*models.User, error)
	StoreRememberMeToken(userId uint, token string) error
	CheckRememberMeToken(token string) error
	DestroyRememberMeToken(userId uint) error
}

type SessionUsecase struct {
	env *env.Env
	userStore storage.UserStore
	sessionStore storage.SessionStore
}

func NewSessionUsecase(env *env.Env, userStore storage.UserStore, sessionStore storage.SessionStore) SessionUsecaseInterface {
	return &SessionUsecase{env: env, userStore: userStore, sessionStore: sessionStore}
}

func (usecase *SessionUsecase) Create(loginForm models.NewSessionForm) (*models.User, error) {
	user, err := usecase.userStore.FindByEmail(loginForm.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginForm.Password)); err != nil {
		e := models.NewErrBadParam(err, []string{"password"}, []string{"invalid_password"})
		return nil, e
	}

	return user, nil
}

func (usecase *SessionUsecase) StoreRememberMeToken(userId uint, token string) error {
	err := usecase.sessionStore.StoreRememberMeToken(userId, token)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *SessionUsecase) CheckRememberMeToken(token string) error {
	err := usecase.sessionStore.CheckRememberMeToken(token)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *SessionUsecase) DestroyRememberMeToken(userId uint) error {
	err := usecase.sessionStore.DestroyRememberMeToken(userId)
	if err != nil {
		return err
	}

	return nil
}