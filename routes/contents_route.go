package routes

import (
	"github.com/psinthorn/gogolang.co/controllers/contents"
)

func ContentRoutes() {
	router.POST("/contents", contents.Create)
	router.GET("/contents", contents.GetAll)
	router.GET("/contents/content_id", contents.Get)
}
