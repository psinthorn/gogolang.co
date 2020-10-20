package app

import (
	"github.com/psinthorn/gogolang.co/controllers/categories"
	"github.com/psinthorn/gogolang.co/controllers/contents"
	"github.com/psinthorn/gogolang.co/controllers/index"
	"github.com/psinthorn/gogolang.co/controllers/users"
)

func urlsMapping() {
	// router.GET("/ping", controllers.Ping.Pong)
	router.POST("/users", users.Create)
	router.GET("/users/:id", users.Get)
	router.PATCH("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)

	router.POST("/categories", categories.Create)
	router.GET("/categories", categories.GetAll)
	router.GET("/categories/:id", categories.Get)

	router.POST("/contents", contents.Create)
	router.GET("/contents", contents.GetAll)
	router.GET("/contents/:id", contents.Get)

	router.GET("/", index.Welcome)
	router.GET("/about", index.About)
}
