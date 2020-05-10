package storage

import "github.com/jsagl/newsfeed-go-server/app/models"

type ArticleStore interface {
	Index() ([]*models.Article, error)
}

type FavoriteStore interface {
	Index(userId uint) ([]*models.Favorite, error)
	Create(favorite *models.Favorite) (*models.Favorite, error)
	BookmarkedUrls(userId uint) ([]string, error)
	DestroyByTargetUrlAndUserId(targetUrl string, userId uint) error
}

type UserStore interface {
	Create(*models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type SessionStore interface {
	StoreRememberMeToken(userId uint, token string) error
	CheckRememberMeToken(token string) error
	DestroyRememberMeToken(userId uint) error
}