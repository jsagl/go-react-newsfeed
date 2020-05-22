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

// Create a new session when the user signs in
func (handler *SessionHttpHandler) Create(c *gin.Context) {
	var loginForm models.NewSessionForm

	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		// Will be necessary to provide more information to user on input error
		e := models.NewErrBadParam(err, []string{}, []string{})
		HandleErrors(c, e)
		return
	}

	user, err := handler.usecase.Create(loginForm)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	handler.setSessionCookie(c, user.ID)

	if loginForm.RememberMe {
		handler.setRememberMeCookie(c, user.ID)
	}

	c.JSON(http.StatusCreated, user)
}

// Refresh session JWT if user is authenticated (checked in middleware)
func (handler *SessionHttpHandler) Refresh(c *gin.Context) {
	handler.setSessionCookie(c, uint(c.GetInt64("userId")))

	c.AbortWithStatus(http.StatusOK)
}

// Verify the rememberMeToken when the user loads the page. If it exists, is valid and was not invalidated in DB, return a session cookie.
func (handler *SessionHttpHandler) CheckRememberMeToken (c *gin.Context) {
	// if a useId is present, it means that the user sent a sessionToken and is already authenticated
	userId := uint(c.GetInt64("userId"))
	if  userId != 0 {
		handler.setSessionCookie(c, userId)
		c.AbortWithStatus(http.StatusOK)
		return
	}

	jwtKey := []byte(os.Getenv("SECRET_KEY"))

	cookie, err := c.Cookie("rememberMeToken")
	if err != nil || cookie == "" {
		e := models.NewErrUnauthorized(err, []string{"token"}, []string{"missing_remember_me_token"})
		c.AbortWithStatusJSON(http.StatusUnauthorized, e.ServeCustomErr())
		return
	}

	session := &models.Session{}
	tkn, err := jwt.ParseWithClaims(cookie, session, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid || session.StandardClaims.Subject != "remember_me" {
		e := models.NewErrUnauthorized(err, []string{"token"}, []string{"invalid_token"})
		c.AbortWithStatusJSON(http.StatusUnauthorized, e.ServeCustomErr())
		return
	}

	tokenInvalidInDB := handler.usecase.CheckRememberMeToken(cookie)

	if tokenInvalidInDB != nil {
		e := models.NewErrUnauthorized(tokenInvalidInDB, []string{"token"}, []string{"invalid_token"})
		c.AbortWithStatusJSON(http.StatusUnauthorized, e.ServeCustomErr())
		return
	}

	handler.setSessionCookie(c, session.UserID)
	c.AbortWithStatus(http.StatusOK)
}

// Destroy rememberMeToken and sessionToken on user logout
func (handler *SessionHttpHandler) Destroy(c *gin.Context) {
	// Not clean but will do for now. To be done properly when session persistence is implemented

	expirationTime := time.Now().Add(- 10 * time.Minute)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "sessionToken",
		Value:   "",
		Expires: expirationTime,
		SameSite: http.SameSiteStrictMode,
		Secure: os.Getenv("ENVIRONMENT") == "PRODUCTION",
		HttpOnly: true,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "rememberMeToken",
		Value:   "",
		Expires: expirationTime,
		SameSite: http.SameSiteStrictMode,
		Secure: os.Getenv("ENVIRONMENT") == "PRODUCTION",
		HttpOnly: true,
	})

	err := handler.usecase.DestroyRememberMeToken(uint(c.GetInt64("userId")))
	if err != nil {} //TODO: better handler error in this case

	c.AbortWithStatus(http.StatusOK)
}


// Create a session cookie
func (handler *SessionHttpHandler) setSessionCookie(c *gin.Context, userId uint) {
	jwtKey := []byte(os.Getenv("SECRET_KEY"))

	expirationTime := time.Now().Add(5 * time.Minute)
	session := &models.Session{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix(), Subject: "session"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "sessionToken",
		Value:   tokenString,
		Expires: expirationTime,
		SameSite: http.SameSiteStrictMode,
		Secure: os.Getenv("ENVIRONMENT") == "PRODUCTION",
		HttpOnly: true,
	})
}

// Create a rememberMeCookie
func (handler *SessionHttpHandler) setRememberMeCookie(c *gin.Context, userId uint) {
	jwtKey := []byte(os.Getenv("SECRET_KEY"))

	expirationTime := time.Now().AddDate(0, 0, 30)
	session := &models.Session{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix(), Subject: "remember_me"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		HandleErrors(c, err)
		return
	}

	e := handler.usecase.StoreRememberMeToken(userId, tokenString)
	if e != nil {
		HandleErrors(c, e)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "rememberMeToken",
		Value:   tokenString,
		Expires: expirationTime,
		SameSite: http.SameSiteStrictMode,
		Secure: os.Getenv("ENVIRONMENT") == "PRODUCTION",
		HttpOnly: true,
	})
}