package app

import (
	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/utils"
)

var (
	router = gin.Default()
)

func StartApp() {
	port := utils.Server.RunningPort()
	urlsMapping()
	router.Run(":" + port)
}
