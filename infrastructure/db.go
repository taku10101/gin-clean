package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDatabase() (*sql.DB, error) {
    connStr := "user=postgres dbname=mydb password=mypassword sslmode=disable port=5411"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    return db, nil
}
