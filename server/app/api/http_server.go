package api

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	v1 "github.com/jsagl/newsfeed-go-server/app/api/v1"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/middleware"
	"github.com/jsagl/newsfeed-go-server/app/usecase"
	"os"
)

func NewHTTPServer(env *env.Env, mw *middleware.Middleware, usecases usecase.Usecases) {
	// Initialize router
	r := gin.New()

	// Call middleware
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	r.Use(gin.Recovery())
	r.Use(mw.Options())
	r.Use(mw.ErrorLogger())
	r.Use(mw.Logging())
	r.Use(mw.HeadersAndCORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	v1ArticleHandler := v1.NewArticleHttpHandler(env, usecases.Articles)
	v1FavoriteHandler := v1.NewFavoriteHttpHandler(env, usecases.Favorites)
	v1UserHandler := v1.NewUserHttpHandler(env, usecases.Users)
	v1SessionHandler := v1.NewSessionHttpHandler(env, usecases.Sessions)

	// V1
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/articles", mw.GetUserIdFromToken(), v1ArticleHandler.Index)

		apiV1.GET("/favorites", mw.VerifyAuthentication(), v1FavoriteHandler.Index)
		apiV1.POST("/favorites", mw.VerifyAuthentication(), v1FavoriteHandler.Create)
		apiV1.DELETE("/favorites", mw.VerifyAuthentication(), v1FavoriteHandler.Destroy)

		apiV1.POST("/signup", v1UserHandler.Create)

		apiV1.POST("/signin", v1SessionHandler.Create)
		apiV1.GET("/refresh", mw.VerifyAuthentication(), v1SessionHandler.Refresh)
		apiV1.GET("/signout", mw.VerifyAuthentication(), v1SessionHandler.Destroy)
	}

	// Start HTTP server
	port := ":" + os.Getenv("PORT")
	r.Run(port)
}