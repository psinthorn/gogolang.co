package app

import (
	controllers "github.com/psinthorn/gogolang.co/controllers/index"
	"github.com/psinthorn/gogolang.co/controllers/ping"
)

func urlsMapping() {
	router.GET("/ping", ping.Pong)
	// router.POST("/users", controllers.Users.CreateUser)
	// router.GET("/users/:user_id", controllers.Users.GetUser)
	router.GET("/", controllers.IndexPage.Welcome)
	router.GET("/about", controllers.IndexPage.About)
}
