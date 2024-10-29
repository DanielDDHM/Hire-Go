# Hire-Go
A platform to connect clients and professionals.

## Project Structure
We will define our initial project structure as follows:

- **cmd**: This will contain our `main.go` files, responsible for starting the application.
- **config**: Configuration files, such as environment variables and logs, will be stored here.
- **internal**: This is where the core business logic will reside.
  - **internal/dto**: Here, we will define the data transfer objects that dictate the allowed data types for input in the application.
  - **internal/handler**: This folder will contain routing files (you can refer to this as controllers if preferred).
  - **internal/database**: Everything related to the database will be stored here.
    - **internal/database/migrations**: This is where our database migration files will go.
    - **internal/database/queries**: SQL queries for database interactions will be stored here.
  - **internal/repository**: This will contain our repository layer.
  - **internal/service**: Finally, our service layer, where the business logic resides (you can refer to this as use cases if preferred).
  - **internal/utils**: This folder will store utility functions, helpers, and shared tools used throughout the application.

## Technologies
- **API Framework**: Gin
- **ORM**: Gorm
- **Database**: PostgreSQL
- **Containerization**: Docker
- **Container Orchestration**: Kubernetes (K8s)
- **Logger**: zerolog
- **API Documentation**: Swaggo
- **Database Migrations**: golang-migrate
- **Validation**: go-playground/validator
