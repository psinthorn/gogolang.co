package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/controllers/index"
)

var (
	router = gin.Default()
)

func IndexRoutes() {
	router.GET("/", index.Welcome)
	router.GET("/about", index.About)
}
