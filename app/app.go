package app

import (
	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/configs"

	"github.com/psinthorn/gogolang.co/logger"
)

var (
	router = gin.Default()
)

func StartApp(port string) {
	router.LoadHTMLGlob("views/*/*.html")
	router.Static("/assets/", "./assets/")
	urlsMapping()

	logger.Info("Start Application...")
	router.Run(":" + configs.ServerPort.PortSelector(port))
}
