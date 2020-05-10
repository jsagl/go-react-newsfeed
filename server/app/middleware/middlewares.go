package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"time"
)

type Middleware struct {
	env *env.Env
}

func InitMiddleware(env *env.Env) *Middleware {
	return &Middleware{env: env}
}

const ReqIdKey = "requestId"


func (mw *Middleware) Options() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			//c.Header("Access-Control-Allow-Origin", os.Getenv("CORS_ALLOW_ORIGIN"))
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, X-Requested-With")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Content-Type", "application/json")
			c.AbortWithStatus(http.StatusOK)
		}
	}
}


func (mw *Middleware) Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		reqId := c.Request.Header.Get("X-Request-ID")
		if reqId == "" {
			reqId = uuid.NewV4().String()
		}
		c.Writer.Header().Set("X-Request-ID", reqId)

		c.Set("requestId", reqId)

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		status := c.Writer.Status()

		mw.env.Logger.Infow("client_request",
			"uuid", reqId,
			"method", c.Request.Method,
			"route", c.Request.RequestURI,
			"remote_address", c.Request.RemoteAddr,
			"host", c.Request.Host,
			"latency", latency,
			"status", status,
		)
	}
}

func (mw *Middleware) ErrorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		reqId, _ := c.Get("requestId")
		for _, e := range c.Errors {
			mw.env.Logger.Errorw("internal_server_error", "uuid", reqId, "error", e)
		}
	}
}

func (mw *Middleware) HeadersAndCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		//c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ALLOW_ORIGIN"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

func (mw *Middleware) VerifyAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtKey := []byte(os.Getenv("SECRET_KEY"))

		cookie, err := c.Cookie("sessionToken")
		if err != nil || cookie == "" {
			e := models.NewErrUnauthorized(err, []string{"token"}, []string{"missing_auth_token"})
			c.AbortWithStatusJSON(http.StatusUnauthorized, e.ServeCustomErr())
		}

		session := &models.Session{}
		tkn, err := jwt.ParseWithClaims(cookie, session, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !tkn.Valid || session.StandardClaims.Subject != "session" {
			e := models.NewErrUnauthorized(err, []string{"token"}, []string{"invalid_token"})
			c.AbortWithStatusJSON(http.StatusUnauthorized, e.ServeCustomErr())
		}

		c.Set("userId", int64(session.UserID))
		c.Next()
	}
}

func (mw *Middleware) GetUserIdFromToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtKey := []byte(os.Getenv("SECRET_KEY"))

		cookie, err := c.Cookie("sessionToken")
		if err != nil || cookie == "" {
			c.Next()
		}

		session := &models.Session{}

		tkn, err := jwt.ParseWithClaims(cookie, session, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !tkn.Valid {
			c.Next()
		}

		c.Set("userId", int64(session.UserID))
		c.Next()
	}
}