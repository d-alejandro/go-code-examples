# Go Code Examples

## Topics

- Back-end
- Clean Architecture
- CRUD
- Docker Compose (Docker 27.3.1, Docker Compose v2.29.7)
- Dockerfile with multi-stage builds
- DTO
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- Golang 1.23
- [Golangci-lint](https://github.com/golangci/golangci-lint)
- [Goose](https://github.com/pressly/goose)
- [GORM](https://github.com/go-gorm/gorm)
- HTTP Server
- Linters
- Linux
- [Logrus](https://github.com/sirupsen/logrus)
- Makefile
- Manual Dependency Injection
- MVC
- PostgreSQL 17
- [Pre-commit](https://github.com/pre-commit/pre-commit)
- [Pre-commit hooks](https://github.com/pre-commit/pre-commit-hooks)
- RESTful API
- SOLID
- SQL

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

4. Build and run the application:

```
make build-run
```

## API Endpoints

### All Orders with pagination

- Request URL: `http://localhost:8081/api/orders?start=0&end=1&sort_column=id&sort_type=asc`
- Method: `GET`
- Path: `/orders`
- Headers: `Accept:application/json, Content-Type:application/json`
- Parameters: `start, end, sort_column, sort_type`
- Status Code: `200`
- Response:

```json
{
  "data": [
    {
      "id": 10000001,
      "agency_name": "ОАО ITНефтьРыбОпт",
      "status": "waiting",
      "is_confirmed": false,
      "is_checked": false,
      "rental_date": "31-12-2023",
      "user_name": "Валерия Фёдоровна Вишнякова",
      "transport_count": 3,
      "guest_count": 3,
      "admin_note": null
    }
  ]
}
```

### Create Order

- Request URL: `http://localhost:8081/api/orders`
- Method: `POST`
- Path: `/orders`
- Headers: `Accept:application/json, Content-Type:application/json`
- Parameters: `agency_name, rental_date, guest_count, transport_count, user_name, email, phone`
- Status Code: `200`
- Response:

```json
{
  "data": {
    "id": 10000011,
    "agency_name": "Test Agency Name",
    "status": "waiting",
    "is_confirmed": false,
    "is_checked": false,
    "rental_date": "01-02-2024",
    "user_name": "Test User Name",
    "transport_count": 1,
    "guest_count": 1,
    "admin_note": "",
    "note": "",
    "email": "test@test.ru",
    "phone": "71437854547",
    "confirmed_at": null,
    "created_at": "31-01-2024 22:11:14",
    "updated_at": "31-01-2024 22:11:14"
  }
}
```

### Order Details

- Request URL: `http://localhost:8081/api/orders/10000011`
- Method: `GET`
- Path: `/orders/{id}`
- Headers: `Accept:application/json, Content-Type:application/json`
- Status Code: `200`
- Response:

```json
{
  "data": {
    "id": 10000011,
    "agency_name": "Test Agency Name",
    "status": "waiting",
    "is_confirmed": false,
    "is_checked": false,
    "rental_date": "01-02-2024",
    "user_name": "Test User Name",
    "transport_count": 1,
    "guest_count": 1,
    "admin_note": "",
    "note": "",
    "email": "test@test.ru",
    "phone": "71437854547",
    "confirmed_at": null,
    "created_at": "31-01-2024 22:11:15",
    "updated_at": "31-01-2024 22:11:15"
  }
}
```

### Update Order

- Request URL: `http://localhost:8081/api/orders/10000011`
- Method: `PUT`
- Path: `/orders/{id}`
- Headers: `Accept:application/json, Content-Type:application/json`
- Parameters: `guest_count, transport_count, user_name, email, phone`
- Status Code: `200`
- Response:

```json
{
  "data": {
    "id": 10000011,
    "agency_name": "Test Agency Name",
    "status": "waiting",
    "is_confirmed": false,
    "is_checked": false,
    "rental_date": "01-02-2024",
    "user_name": "Test7",
    "transport_count": 7,
    "guest_count": 7,
    "admin_note": "qq",
    "note": "ww",
    "email": "777@777.test",
    "phone": "71111111111",
    "confirmed_at": null,
    "created_at": "31-01-2024 22:11:15",
    "updated_at": "31-01-2024 22:17:47"
  }
}
```

### Delete Order

- Request URL: `http://localhost:8081/api/orders/10000011`
- Method: `DELETE`
- Path: `/orders/{id}`
- Headers: `Accept:application/json, Content-Type:application/json`
- Status Code: `200`
- Response:

```json
{
  "data": {
    "id": 10000011,
    "agency_name": "Test Agency Name",
    "status": "waiting",
    "is_confirmed": false,
    "is_checked": false,
    "rental_date": "01-02-2024",
    "user_name": "Test7",
    "transport_count": 7,
    "guest_count": 7,
    "admin_note": "qq",
    "note": "ww",
    "email": "777@777.test",
    "phone": "71111111111",
    "confirmed_at": null,
    "created_at": "31-01-2024 22:11:15",
    "updated_at": "31-01-2024 22:17:47"
  }
}
```
