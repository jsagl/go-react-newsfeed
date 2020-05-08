package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/usecase"
	"net/http"
)

type FavoriteHttpHandler struct {
	usecase usecase.FavoriteUsecaseInterface
	env  *env.Env
}

func NewFavoriteHttpHandler(env *env.Env, usecase usecase.FavoriteUsecaseInterface) *FavoriteHttpHandler {
	return &FavoriteHttpHandler{usecase: usecase, env: env}
}

func (handler *FavoriteHttpHandler) Index(c *gin.Context) {
	favorites, err := handler.usecase.Index(c)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, favorites)
}

func (handler *FavoriteHttpHandler) Create(c *gin.Context) {
	favorite, err := handler.usecase.Create(c)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	c.JSON(http.StatusCreated, favorite)
}

func (handler *FavoriteHttpHandler) Destroy(c *gin.Context) {
	err := handler.usecase.Destroy(c)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	c.JSON(http.StatusNoContent, "")
}