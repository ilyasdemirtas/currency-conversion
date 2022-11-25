# Currency Conversion API

## Installation

The MySQL database should create on your local machine. DB and other information can change with .env file.
- MySQL tables will be ready when project run by migration of the [Gorm package](https://gorm.io/)
- Default data will be insert when project run.
- All currencies with the code stored at the `currencies` table
- Default credentials are U: ilyasdemirtas@hotmail.com.tr P: 123qwe

### Docker:
In the root dir of project: `docker-compose up`

### Manuel:
In the root dir of project: `go run .`

## API Documantation:

Documentations of the endpoints are on the Swagger UI.

Base URL: `localhost:8080/api`

Swagger URL: `localhost:8080/docs`
