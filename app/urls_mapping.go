package app

import (
	"github.com/psinthorn/gogolang.co/controllers/index"

	"github.com/psinthorn/gogolang.co/controllers/users"
)

func urlsMapping() {
	// router.GET("/ping", controllers.Ping.Pong)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

	router.GET("/", index.Welcome)
	router.GET("/about", index.About)
}
