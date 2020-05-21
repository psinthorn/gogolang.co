package app

import (
	"github.com/psinthorn/gogolang.co/controllers/ping"
	controllers "github.com/psinthorn/gogolang.co/controllers/users"
)

func urlsMapping() {
	router.GET("/ping", ping.Pong)
	router.POST("/users", controllers.Users.CreateUser)
	router.GET("/users/:user_id", controllers.Users.GetUser)
}
