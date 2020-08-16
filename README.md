# env: Parse environment variables, map to structs

![Go](https://github.com/aacanakin/env/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/aacanakin/env)](https://goreportcard.com/report/github.com/aacanakin/env) [![Coverage Status](https://coveralls.io/repos/github/aacanakin/env/badge.svg?branch=master)](https://coveralls.io/github/aacanakin/env?branch=master)

env is a mapper from environment variables to structs

## Features

- Ability to map environment variables to structs
- Ability to customize environment variables with struct tags
- Ability to parse nested/sub structs
- Ability to parse embedded nested/sub structs
- Ability to parse and map types (`bool`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, `complex64`, `complex128`, `struct`)

## Getting started

### Install

```sh
// Requires go1.15+

go get github.com/aacanakin/env
```

### Usage

```go
// main.go
package main

import "fmt"
import "github.com/aacanakin/env"

type Config struct {
  Host string
  Port int
  Debug bool
}

func main() {

  var c Config

  err := env.Parse(&c)
  if err != nil {
    panic(err)
  }

  fmt.Println("Host: ", c.Host)
  fmt.Println("Port: ", c.Port)
  fmt.Println("Debug: ", c.Debug)
}
```

Run

```sh
$ HOST=localhost PORT=8081 debug=true go run main.go

// Output
Host:  localhost
Port:  8081
Debug:  true
```

### Custom environment var mapping

```go
// main.go
package main

import "fmt"
import "github.com/aacanakin/env"

func main() {
  type Config struct {
    Host string `env:"SERVICE_HOST"`
    Port int    `env:"SERVICE_PORT"`
    Debug bool  `env:"SERVICE_DEBUG"`
  }

  var c Config

  err := env.Parse(&c)
  if err != nil {
    panic(err)
  }

  fmt.Println("Host: ", c.Host)
  fmt.Println("Port: ", c.Port)
  fmt.Println("Debug: ", c.Debug)
}
```

Run

```sh
SERVICE_HOST=localhost SERVICE_PORT=8081 SERVICE_DEBUG=true go run main.go

// Output
Host:  localhost
Port:  8081
Debug:  true
```

### Nested/Sub structs

```go
// main.go
package main

import "fmt"
import "github.com/aacanakin/env"

type db struct {
  Host string `env:"DB_HOST"`
  Port int `env:"DB_PORT"`
}

type service struct {
  Debug bool `env:"SERVICE_DEBUG"`
}

type Config struct {
  DB db
  Service service
}

func main() {
  var c Config

  err := env.Parse(&c)
  if err != nil {
    panic(err)
  }

  fmt.Println("DB Host: ", c.DB.Host)
  fmt.Println("DB Port: ", c.DB.Port)
  fmt.Println("Service Debug: ", c.Service.Debug)
}

```

Run

```sh
DB_HOST=localhost DB_PORT=3306 SERVICE_DEBUG=true go run main.go

// Output
DB Host:  localhost
DB Port:  3306
Service Debug:  true
```

### Using with .env

Install

```sh
go get github.com/joho/godotenv
```

Create a .env file

```sh
# .env
HOST=localhost
PORT=8081
```

```go
// main.go
package main

import (
  "github.com/joho/godotenv"
  "github.com/aacanakin/env"
  "log"
)

type Config struct {
  Host string
  Port uint16
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  var c Config

  err = env.Parse(&c)

  log.Println("Host: ", c.Host)
  log.Println("Port: ", c.Port)
}
```

### Readonly config

Since it's a good practice to keep environment variables immutable, it will be also a good practice to keep the parsed struct immutable too. However, Go's reflect package doesn't allow setting unexported fields. So,
to keep the struct immutable;

- Create a config package

```go
package config

import "github.com/aacanakin/env"

type conf struct {
	Host string
	Port uint16
}

type Config struct {
	conf conf
}

func (c Config) Host() string {
	return c.conf.Host
}

func (c Config) Port() uint16 {
	return c.conf.Port
}

func New() (*Config, error) {
	var c conf
	err := env.Parse(&c)
	if err != nil {
		return nil, err
	}

	return &Config{c}, nil
}
```

- Use it in your main.go

```go
package main

import (
	"fmt"

	"github.com/aacanakin/env_test/config"
)

func main() {
	c, err := config.New()
	if err != nil {
		panic(err)
	}

  // HERE, c.conf is not accessible
	fmt.Println("Host:", c.Host())
	fmt.Println("Port:", c.Port())
}

```

This looks like it's not very idiomatic go. Feedbacks here are welcome.

## Roadmap

- [ ] `omitempty` tag constraint
- [ ] `file` tag constraint
- [ ] Ability to customize tag key (default is env)
- [ ] Provide an instance based parser
- [ ] Provide a read only config example with custom conf package
- [ ] Release v0.1
