package controllers

import (
	"net/http"
	"strconv"

	"github.com/taku10101/gin-clean/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    UserUsecase *usecases.UserUsecase
}

func (uc *UserController) GetUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := uc.UserUsecase.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}
