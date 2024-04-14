package infrastructure

import (
	"database/sql"

	"github.com/taku10101/gin-clean/entities"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) FindByID(id int) (*entities.User, error) {
    query := "SELECT id, name, age FROM users WHERE id = $1"
    row := r.DB.QueryRow(query, id)

    var user entities.User
    if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
        return nil, err
    }
    return &user, nil
}
