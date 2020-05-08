package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/usecase"
	"net/http"
)

type ArticleHttpHandler struct {
	usecase usecase.ArticleUsecaseInterface
	env  *env.Env
}

func NewArticleHttpHandler(env *env.Env, usecase usecase.ArticleUsecaseInterface) *ArticleHttpHandler {
	return &ArticleHttpHandler{usecase: usecase, env: env}
}

func (handler *ArticleHttpHandler) Index(c *gin.Context) {
	articles, err := handler.usecase.Index(c)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, articles)
}
