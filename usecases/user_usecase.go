package usecases

import (
	"github.com/taku10101/gin-clean/entities"
	"github.com/taku10101/gin-clean/interfaces"
)

type UserUsecase struct {
    UserRepo interfaces.UserRepository
}

func (u *UserUsecase) GetUser(id int) (*entities.User, error) {
    return u.UserRepo.FindByID(id)
}
