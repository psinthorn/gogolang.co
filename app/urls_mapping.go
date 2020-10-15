package app

import (
	"github.com/psinthorn/gogolang.co/routes"
)

func urlsMapping() {
	// router.GET("/ping", controllers.Ping.Pong)
	routes.UserRoutes()
	routes.ContentRoutes()
	routes.IndexRoutes()
}
