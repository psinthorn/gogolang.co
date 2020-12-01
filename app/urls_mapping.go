package app

import (
	"github.com/psinthorn/gogolang.co/controllers/categories"
	"github.com/psinthorn/gogolang.co/controllers/contents"
	"github.com/psinthorn/gogolang.co/controllers/index"
	"github.com/psinthorn/gogolang.co/controllers/users"
)

func urlsMapping() {

	// Section: Index
	router.GET("/", index.Welcome)
	router.GET("/about", index.About)
	// router.GET("/users", users.GetAll)
	// router.GET("/contents", contents.GetAll)

	// Index: API
	// router.GET("/ping", controllers.Ping.Pong)
	router.POST("/users", users.Create)
	router.GET("/users", users.GetAll)
	router.GET("/users/:id", users.Get)
	router.PATCH("/users/:id", users.Update)
	router.PUT("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)

	router.POST("/categories", categories.Create)
	router.GET("/categories", categories.GetAll)
	router.GET("/categories/:id", categories.Get)

	router.POST("/contents", contents.Create)
	router.GET("/contents", contents.GetAll)
	router.GET("/contents/:id", contents.Get)
	router.PATCH("/contents/:id", contents.Update)
	router.PUT("/contents/:id", contents.Update)
	router.DELETE("/contents/:id", contents.Delete)

}
