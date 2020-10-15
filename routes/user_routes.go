package routes

import "github.com/psinthorn/gogolang.co/controllers/users"

func UserRoutes() {
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

}
