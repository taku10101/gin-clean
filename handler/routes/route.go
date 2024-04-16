package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taku10101/gin-clean/controllers"
)

type Route struct {
	userController *controllers.UserController
}


func NewRoute(userController *controllers.UserController) *Route {
	return &Route{userController: userController}
}

func Routes(r *Route) {
	router := gin.Default()
	router.GET("/users", r.userController.Index)
	router.GET("/users/:id", r.userController.Show)
	router.POST("/users", r.userController.Create)
	router.PUT("/users/:id", r.userController.Update)
	router.DELETE("/users/:id", r.userController.Delete)
	
	router.Run(":8080")
}