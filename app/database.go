package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	User   string
	Passwd string
	Net    string
	Addr   string
	DbName string
	MySql  mysql.Config
}

func defaultDbConfig() DbConfig {
	return DbConfig{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DbName: "gym_log",
	}

}

type DbConfigFunc func(*DbConfig)

func newDbConfig(funcs ...DbConfigFunc) DbConfig {
	d := defaultDbConfig()
	for _, fn := range funcs {
		fn(&d)
	}
	d.MySql = mysql.Config{
		User:   d.User,
		Passwd: d.Passwd,
		Net:    d.Net,
		Addr:   d.Addr,
		DBName: d.DbName,
	}
	return d
}

func withDbUser(usr string) DbConfigFunc {
	return func(d *DbConfig) {
		d.User = usr
	}
}

func withDbPass(pass string) DbConfigFunc {
	return func(d *DbConfig) {
		d.Passwd = pass
	}
}

func withDbNet(net string) DbConfigFunc {
	return func(d *DbConfig) {
		d.Net = net
	}
}

func withDbAddr(addr string) DbConfigFunc {
	return func(d *DbConfig) {
		d.Addr = addr
	}
}

func withDbName(name string) DbConfigFunc {
	return func(d *DbConfig) {
		d.DbName = name
	}
}

func LoadDb(db *sql.DB) error {
	cfg := defaultDbConfig()

	var err error
	db, err = sql.Open("mysql", cfg.MySql.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	fmt.Println("Connected to the database!")
	return nil
}
