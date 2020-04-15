package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
