# Go Code Examples

## Topics

- Back-end
- Clean Architecture
- CRUD
- Debugging Mode
- [Delve](https://github.com/go-delve/delve)
- Dependency Injection
- Design Patterns
- Docker Compose
- Dockerfile with multi-stage builds
- DTO
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- Golang 1.21
- [Goose](https://github.com/pressly/goose)
- [GORM](https://github.com/go-gorm/gorm)
- Linux
- [Logrus](https://github.com/sirupsen/logrus)
- Makefile
- MVC
- PostgreSQL 16
- RESTful API
- SOLID
- SQL
- [Viper](https://github.com/spf13/viper)

## Installation

1. Clone this repository:

```
git clone git@github.com:d-alejandro/go-code-examples.git
```

2. Go to `go-code-examples` directory:

```
cd go-code-examples
```

3. Create a new .env file:

```
cp .env.example .env
```

4. Download and install `Go` SDK `1.21.6`.

5. Build and run the application:

```
make run
```

or

```
make
```

6. Start the debugging process in the IDE with `Go Remote` configuration.

```
Host: localhost
Port: 7000
```

7. Apply migrations:

```
make migrate
```
