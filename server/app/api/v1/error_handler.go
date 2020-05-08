package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"net/http"
)

func HandleErrors(c *gin.Context, err error) {
	requestId := c.GetString("requestId")

	if e, ok := err.(*models.ErrBadParam); ok {
		c.JSON(http.StatusBadRequest, e.ServeCustomErr())
		return
	}

	if e, ok := err.(*models.ErrUnauthorized); ok {
		c.JSON(http.StatusUnauthorized, e.ServeCustomErr())
		return
	}

	if e, ok := err.(*models.ErrNotFound); ok {
		c.JSON(http.StatusNotFound, e.ServeCustomErr())
		return
	}

	e := models.NewErrInternal(requestId, err)
	c.Error(err)
	c.JSON(http.StatusInternalServerError, e.ServeCustomErr())
}