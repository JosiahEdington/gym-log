package data

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/JosiahEdington/gym-log/app"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var gymDB *GymDB

type GymDB struct {
	db *sql.DB
}

func ConnectToDB(cfg *app.Config) error {
	//TODO: Add a way to store these as secrets.
	var (
		dbUser                 = cfg.DB.User
		dbPwd                  = cfg.DB.Passwd
		dbName                 = cfg.DB.DBName
		instanceConnectionName = cfg.Instance
		// dbUser                 = mustGetEnv("gym-log-backend")
		// dbPwd                  = mustGetEnv("+-*qq9kUSg/X]OU#")
		// dbName                 = mustGetEnv("gym-log-01")
		// instanceConnectionName = mustGetEnv("gym-log-app-448001:us-west4:gym-log-01")
		usePrivate = os.Getenv("PRIVATE_IP")
	)

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}

	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true", dbUser, dbPwd, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	db = dbPool
	return nil
}

func GetGymDB() *GymDB {
	return gymDB
}
