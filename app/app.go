package app

import (
	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/configs"
)

var (
	router = gin.Default()
)

func StartApp() {
	router.LoadHTMLGlob("views/*/*.html")
	urlsMapping()
	router.Run(":" + configs.ServerPort.PortSelector("8090"))
}
