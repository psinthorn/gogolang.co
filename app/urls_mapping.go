package app

import (
	controllers "github.com/psinthorn/gogolang.co/controllers/index"
)

func urlsMapping() {
	// router.GET("/ping", ping.Pong)
	// router.POST("/users", controllers.Users.CreateUser)
	// router.GET("/users/:user_id", controllers.Users.GetUser)
	router.GET("/", controllers.IndexPage.Welcome)
}
