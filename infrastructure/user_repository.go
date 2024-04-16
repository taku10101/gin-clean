package infrastructure

import (
	"database/sql"

	"github.com/taku10101/gin-clean/entities"
)

type UserRepository struct {
    DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]*entities.Users, error) {
    var users []*entities.Users
    query := `SELECT id, name, age FROM users`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        user := &entities.Users{}
        if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}
func (r *UserRepository) Add(user *entities.Users) error {
    query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
    err := r.DB.QueryRow(query, user.Name, user.Age).Scan(&user.ID)//Scanは一行を取得する
    return err
}

func (r *UserRepository) FindByID(id int) (*entities.Users, error) {
    user := &entities.Users{}
    query := `SELECT id, name, age FROM users WHERE id = $1`
    if err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Age); err != nil {
        return nil, err
    }
    return user, nil
}



func (r *UserRepository) Update(user *entities.Users) error {
    query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`
    _, err := r.DB.Exec(query, user.Name, user.Age, user.ID)//Execは複数行を取得する
    return err
}

func (r *UserRepository) Delete(id int) error {
    query := `DELETE FROM users WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    return err
}


