package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
)

type FavoriteUsecaseInterface interface {
	Index(c *gin.Context) ([]*models.Favorite, error)
	Create(c *gin.Context) (*models.Favorite, error)
	Destroy(c *gin.Context) error
}

type FavoriteUsecase struct {
	env *env.Env
	store storage.FavoriteStore
}

func NewFavoriteUsecase(env *env.Env, db storage.FavoriteStore) FavoriteUsecaseInterface {
	return &FavoriteUsecase{env: env, store: db}
}

func (usecase *FavoriteUsecase) Index(c *gin.Context) ([]*models.Favorite, error) {
	userId := uint(c.GetInt64("userId"))

	favorites, err := usecase.store.Index(userId)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func (usecase *FavoriteUsecase) Create(c *gin.Context) (*models.Favorite, error) {
	var input models.FavoriteInput

	// Will be necessary to better handle this error as for now in case of invalid input,
	// an internal_server_error is returned by the error handler
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	fav := models.ConvertToFavorite(&input)
	fav.UserID = uint(c.GetInt64("userId"))

	favorite, err := usecase.store.Create(fav)
	if err != nil {
		return nil, err
	}

	return favorite, nil
}

func (usecase *FavoriteUsecase) Destroy(c *gin.Context) error {
	var input models.FavoriteDestroy

	// Will be necessary to better handle this error as for now in case of invalid input,
	// an internal_server_error is returned by the error handler
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	err := usecase.store.DestroyByTargetUrlAndUserId(*input.TargetUrl, uint(c.GetInt64("userId")))
	if err != nil {
		return err
	}

	return nil
}