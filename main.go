package main

import (
	"github.com/taku10101/gin-clean/controllers"
	"github.com/taku10101/gin-clean/infrastructure"
	"github.com/taku10101/gin-clean/usecases"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    db, err := infrastructure.NewDatabase()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    userRepo := &infrastructure.UserRepository{DB: db}
    userUsecase := &usecases.UserUsecase{UserRepo: userRepo}
    userController := &controllers.UserController{UserUsecase: userUsecase}

    r := gin.Default()
    r.GET("/users/:id", userController.GetUser)

    r.Run() // listen and serve on 0.0.0.0:8080
}
