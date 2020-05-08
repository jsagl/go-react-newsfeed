package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/usecase"
	"net/http"
)

type UserHttpHandler struct {
	usecase usecase.UserUsecaseInterface
	env  *env.Env
}

func NewUserHttpHandler(env *env.Env, usecase usecase.UserUsecaseInterface) *UserHttpHandler {
	return &UserHttpHandler{usecase: usecase, env: env}
}

func (handler *UserHttpHandler) Create(c *gin.Context) {
	user, err := handler.usecase.Create(c)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}