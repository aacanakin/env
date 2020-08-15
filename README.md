# env: Simple environment variable parser/mapper

env is a mapper for environment variables to structs

## Features

- Ability to map environment variables to structs
- Ability to customize environment variables with struct tags
- Ability to parse nested/sub structs
- Ability to parse embedded nested/sub structs
- Ability to parse and map types (`bool`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, `complex64`, `complex128`, `struct`)

## Getting started

### Install

```sh
go get github.com/aacanakin/env
```

### Usage

```go
// main.go
package main

import "fmt"
import "github.com/aacanakin/env"

func main() {
  type Config struct {
    Host string
    Port int
    Debug bool
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

Run;

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
DB Port:  3345
Service Debug:  true
```
