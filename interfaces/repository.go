package interfaces

import "github.com/taku10101/gin-clean/entities"


type UserRepository interface {
	FindAll() ([]*entities.Users, error)
    Add(user *entities.Users) error 
	FindByID(id int) (*entities.Users, error)
	Update(user *entities.Users) error
	Delete(id int) error
}

