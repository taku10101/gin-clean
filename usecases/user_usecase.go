package usecases

import (
	"github.com/taku10101/gin-clean/entities"
	"github.com/taku10101/gin-clean/interfaces"
)

type UserUsecase struct {
    UserRepo interfaces.UserRepository
}

func NewUserUsecase(userRepo interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (u *UserUsecase) FindAllUser() ([]*entities.Users, error) {
	return u.UserRepo.FindAll()
}

func (u *UserUsecase) FindByIDUser(id int) (*entities.Users, error) {
	return u.UserRepo.FindByID(id)
}


func (u *UserUsecase) AddUser(user *entities.Users) error {
	return u.UserRepo.Add(user)
}

func (u *UserUsecase) UpdateUser(user *entities.Users) error {
	return u.UserRepo.Update(user)
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.UserRepo.Delete(id)
}
