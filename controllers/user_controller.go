package controllers

import (
	"net/http"
	"strconv"

	"github.com/taku10101/gin-clean/entities"
	"github.com/taku10101/gin-clean/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    UserUsecase *usecases.UserUsecase
}

func NewUserController(u *usecases.UserUsecase) *UserController {
    return &UserController{UserUsecase: u}
}

func (c *UserController) Index(ctx *gin.Context) {

    users, err := c.UserUsecase.FindAllUser()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *UserController) Show(ctx *gin.Context) {
    id := ctx.Param("id")
    idInt, err := strconv.Atoi(id)
    user, err := c.UserUsecase.FindByIDUser(idInt)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c *UserController) Create(ctx *gin.Context) {

    user := &entities.Users{}
    if err := ctx.Bind(user); err != nil {//Bindはリクエストの値を構造体にバインドする
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.UserUsecase.AddUser(user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"user": user})
}



func (c *UserController) Update(ctx *gin.Context) {
    id := ctx.Param("id")
    idInt , err := strconv.Atoi(id)
    user, err := c.UserUsecase.FindByIDUser(idInt)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := ctx.Bind(user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.UserUsecase.UpdateUser(user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c *UserController) Delete(ctx *gin.Context) {
    id := ctx.Param("id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.UserUsecase.DeleteUser(idInt); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusNoContent, gin.H{})
}






