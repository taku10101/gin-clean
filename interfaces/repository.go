package interfaces

import "github.com/taku10101/gin-clean/entities"
type UserRepository interface {
    Add(user *entities.User) error //*はポインタ型
	FindAll() ([]*entities.User, error)
 	FindByID(id int) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id int) error

	
}
