package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string       `json:"host"`
	Port     string       `json:"port"`
	Instance string       `json:"instance"`
	DB       mysql.Config `json:"mysql"`
}

func newConfig(funcs ...ConfigFunc) Config {
	c := defaultConfig()
	for _, fn := range funcs {
		fn(&c)
	}
	return c
}

func newDbConfig(funcs ...DbConfigFunc) mysql.Config {
	d := defaultDbConfig()
	for _, fn := range funcs {
		fn(&d)
	}
	return d
}

func LoadConfig() Config {
	var (
		p, _ = os.Getwd()
		need = false
		c    = newConfig()
	)

	file, err := os.OpenFile(p+"/config/config.json", os.O_RDWR, 0777)
	if err == os.ErrExist {
		need = true
	} else if err != nil {
		need = true
		fmt.Printf("error opening file: %v", err)
	}

	if need {
		fileEncoder[Config](file, c)
		return c
	} else {
		c, _ = fileDecoder[Config](file)
	}
	// fileEncoder[Config](file, c)
	return c
}

func defaultConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     "3000",
		Instance: "",
		// Instance: "gym-log-app-448001:us-west4:gym-log-01",
		DB: defaultDbConfig(),
	}
}
func defaultDbConfig() mysql.Config {
	return mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "126.0.0.1:3000",
		DBName:    "gym_log",
		ParseTime: true,
	}
}

type ConfigFunc func(*Config)
type DbConfigFunc func(*mysql.Config)

func withPort(port string) ConfigFunc {
	return func(c *Config) {
		c.Port = port
	}
}

func withHost(host string) ConfigFunc {
	return func(c *Config) {
		c.Host = host
	}
}

func withDbConfig(db mysql.Config) ConfigFunc {
	return func(c *Config) {
		c.DB = db
	}
}

func withDbUser(usr string) DbConfigFunc {
	return func(d *mysql.Config) {
		d.User = usr
	}
}

func withDbPass(pass string) DbConfigFunc {
	return func(d *mysql.Config) {
		d.Passwd = pass
	}
}

func withDbNet(net string) DbConfigFunc {
	return func(d *mysql.Config) {
		d.Net = net
	}
}

func withDbAddr(addr string) DbConfigFunc {
	return func(d *mysql.Config) {
		d.Addr = addr
	}
}

func withDbName(name string) DbConfigFunc {
	return func(d *mysql.Config) {
		d.DBName = name
	}
}

func fileEncoder[T any](file *os.File, v T) error {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(v)
	if err != nil {
		return fmt.Errorf("encode json to file: %w", err)
	}
	return nil
}

func fileDecoder[T any](file *os.File) (T, error) {
	decoder := json.NewDecoder(file)
	var v T

	err := decoder.Decode(&v)
	if err != nil {
		return v, fmt.Errorf("decode json from file: %w", err)
	}
	return v, nil
}
