package main

import (
	"fmt"

	"github.com/aacanakin/env"
)

func main() {
	type db struct {
		Host string `env:"DB_HOST"`
		Port uint16 `env:"DB_PORT"`
	}

	type service struct {
		Debug bool `env:"SERVICE_DEBUG"`
	}

	type Config struct {
		DB      db
		Service service
	}

	var c Config

	err := env.Parse(&c)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("DB Host: ", c.DB.Host)
	fmt.Println("DB Port: ", c.DB.Port)
	fmt.Println("Service Debug: ", c.Service.Debug)
}
