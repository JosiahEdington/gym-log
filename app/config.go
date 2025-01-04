package app

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func defaultConfig() Config {
	return Config{
		Host: "localhost",
		Port: "3000",
	}
}

type ConfigFunc func(*Config)

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

func newConfig(funcs ...ConfigFunc) Config {
	c := defaultConfig()
	for _, fn := range funcs {
		fn(&c)
	}
	return c
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
		tmpC, err := fileDecoder[Config](file)
		if err == nil {
			c = tmpC
		}
	}
	fmt.Printf("config loaded: %v\n", c)
	return c
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
