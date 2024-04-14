package interfaces

import "github.com/taku10101/gin-clean/entities"
type UserRepository interface {
    FindByID(id int) (*entities.User, error)
}
