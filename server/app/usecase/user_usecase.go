package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseInterface interface {
	Create(c *gin.Context) (*models.User, error)
}

type UserUsecase struct {
	env *env.Env
	store storage.UserStore
}

func NewUserUsecase(env *env.Env, db storage.UserStore) UserUsecaseInterface {
	return &UserUsecase{env: env, store: db}
}

func (usecase *UserUsecase) Create(c *gin.Context) (*models.User, error) {
	var user models.NewUserForm

	err := c.ShouldBindJSON(&user)
	if err != nil {
		// Will be necessary to provide more information to user on input error
		e := models.NewErrBadParam(err, []string{}, []string{})
		return nil, e
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	createdUser, err := usecase.store.Create(models.ConvertNewUserFormToUser(&user))
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}