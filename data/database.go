package data

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var gymDB *GymDB

type GymDB struct {
	db *sql.DB
}

func newGymDB(new *sql.DB) *GymDB {
	gymDB = &GymDB{
		db: new,
	}

	return gymDB
}

func ConnectToDB(cfg *mysql.Config) error {
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	fmt.Printf("Connected to database: %v\n", cfg.DBName)
	newGymDB(db)
	return nil
}

func GetGymDB() *GymDB {
	return gymDB
}
