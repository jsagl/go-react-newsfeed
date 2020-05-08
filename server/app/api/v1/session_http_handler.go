package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/usecase"
	"net/http"
	"os"
	"time"
)

type SessionHttpHandler struct {
	usecase usecase.SessionUsecaseInterface
	env  *env.Env
}

func NewSessionHttpHandler(env *env.Env, usecase usecase.SessionUsecaseInterface) *SessionHttpHandler {
	return &SessionHttpHandler{usecase: usecase, env: env}
}

func (handler *SessionHttpHandler) Create(c *gin.Context) {
	user, err := handler.usecase.Create(c)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	jwtKey := []byte(os.Getenv("SECRET_KEY"))
	expirationTime := time.Now().Add(5 * time.Minute)
	session := &models.Session{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		SameSite: http.SameSiteStrictMode,
		Secure: os.Getenv("ENVIRONMENT") == "PRODUCTION",
		HttpOnly: true,
	})

	c.JSON(http.StatusCreated, user)
}

func (handler *SessionHttpHandler) Refresh(c *gin.Context) {
	jwtKey := []byte(os.Getenv("SECRET_KEY"))

	expirationTime := time.Now().Add(5 * time.Minute)

	session := &models.Session{
		UserID: uint(c.GetInt64("userId")),
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		SameSite: http.SameSiteStrictMode,
		Secure: os.Getenv("ENVIRONMENT") == "PRODUCTION",
		HttpOnly: true,
	})

	c.AbortWithStatus(http.StatusOK)
}

func (handler *SessionHttpHandler) Destroy(c *gin.Context) {
	// Not clean but will do for now. To be done properly when session persistence is implemented

	jwtKey := []byte(os.Getenv("SECRET_KEY"))

	expirationTime := time.Now().Add(10 * time.Millisecond)

	session := &models.Session{
		UserID: uint(c.GetInt64("userId")),
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		SameSite: http.SameSiteNoneMode,
		Secure: true,
		HttpOnly: true,
	})

	c.AbortWithStatus(http.StatusOK)
}