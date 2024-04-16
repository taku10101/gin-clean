package main

import (
	"log"

	"github.com/taku10101/gin-clean/controllers"
	"github.com/taku10101/gin-clean/handler/routes"
	"github.com/taku10101/gin-clean/infrastructure"
	"github.com/taku10101/gin-clean/usecases"
)

func main() {
    

    db, err := infrastructure.NewDatabase()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    UserRepo := infrastructure.NewUserRepository(db)
    userUsecase := usecases.NewUserUsecase(UserRepo)
    userController := controllers.NewUserController(userUsecase)

    route := routes.NewRoute(userController)
    routes.Routes(route)
}
