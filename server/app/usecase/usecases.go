package usecase

type Usecases struct {
	Articles ArticleUsecaseInterface
	Favorites FavoriteUsecaseInterface
	Users UserUsecaseInterface
	Sessions SessionUsecaseInterface
}