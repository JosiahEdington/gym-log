package app

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func LoadDb(cfg *mysql.Config) (*sql.DB, error) {
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return db, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return db, pingErr
	}
	fmt.Println("Connected to the database!")
	return db, nil
}
