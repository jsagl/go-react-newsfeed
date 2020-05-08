package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
)

type ArticleUsecaseInterface interface {
	Index(c *gin.Context) ([]*models.Article, error)
}

type ArticleUsecase struct {
	env *env.Env
	store storage.ArticleStore
	favoriteStore storage.FavoriteStore
}

func NewArticleUsecase(env *env.Env, store storage.ArticleStore, favoriteStore storage.FavoriteStore) ArticleUsecaseInterface {
	return &ArticleUsecase{env: env, store: store, favoriteStore: favoriteStore}
}

func (usecase *ArticleUsecase) Index(c *gin.Context) ([]*models.Article, error) {
	articles, err := usecase.store.Index()
	if err != nil {
		return nil, err
	}

	userId := uint(c.GetInt64("userId"))
	if userId == 0 {
		return articles, nil
	}

	bookmarkedUrls, err := usecase.favoriteStore.BookmarkedUrls(userId)
	if err != nil {
		return articles, nil
	}

	setBookmarkedArticles(bookmarkedUrls, articles)

	return articles, nil
}

func setBookmarkedArticles(bookmarks []string, articles []*models.Article) {
	for _, article := range articles {
		articleTargetUrl := article.TargetUrl
		if contains(bookmarks, articleTargetUrl) {
			article.Bookmarked = true
		}
	}
}

func contains(slice []string, str string) bool {
	for _, elem := range slice {
		if elem == str {
			return true
		}
	}
	return false
}