package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jsagl/newsfeed-go-server/app/api"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/middleware"
	"github.com/jsagl/newsfeed-go-server/app/storage/postgresql"
	"github.com/jsagl/newsfeed-go-server/app/storage/scrapped"
	"github.com/jsagl/newsfeed-go-server/app/usecase"
	"log"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("could not load environment variables")
	}

	// Setup Env (logger & timeout)
	env, err := env.NewEnv()
	if err != nil {
		log.Fatalf("Failed to setup env: %v", err)
		return
	}
	defer env.Logger.Sync()

	// Initialize database
	fmt.Println("Initializing postgres database...")
	db, err := postgresql.NewPostgresDatabase()
	if err != nil {
		env.Logger.Fatalw("Failed to initialize database", "error", err)
		return
	}
	defer db.Close()

	// Initialize stores
	fmt.Println("Initializing stores...")
	articleStore := scrapped.NewExternalArticleStore(env)
	favoriteStore := postgresql.NewPostgresFavoriteStore(env, db)
	userStore := postgresql.NewPostgresUserStore(env, db)

	// Initialize usecases
	fmt.Println("Initializing usecases...")
	articleUsecase := usecase.NewArticleUsecase(env, articleStore, favoriteStore)
	favoriteUsecase := usecase.NewFavoriteUsecase(env, favoriteStore)
	userUsecase := usecase.NewUserUsecase(env, userStore)
	sessionUsecase := usecase.NewSessionUsecase(env, userStore)

	usecases := usecase.Usecases{
		Articles: articleUsecase, Favorites: favoriteUsecase,
		Users: userUsecase, Sessions: sessionUsecase,
	}

	// Initialize middleware
	mw := middleware.InitMiddleware(env)

	// Initialize http server
	fmt.Println("Initializing server...")
	api.NewHTTPServer(env, mw, usecases)

}