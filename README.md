Leal Technical Test
This repository contains a Go project for managing an accumulated rewards system with token authentication and multiple endpoints support.

Requirements
Go 1.22 or higher
Docker and Docker Compose
PostgreSQL
Installing Dependencies
To install the necessary project dependencies, make sure to run the following command from the root of the project:

go mod tidy
This command will download and install all the required packages for the project.

Running the Project
There are two ways to run the project:

1. Running Directly from the Source Code
If you want to run the project directly from the source code, follow these steps:

Make sure you are in the root directory of the project.

Run the following command to start the application:

go run cmd/main.go
This will start the Go server on the port configured in the environment file.

2. Running with Docker Compose
If you prefer to use Docker Compose to manage both the PostgreSQL database and the Go application, follow these steps:

Navigate to the compose folder within the project:

cd compose
Run the following command to start the containers in the background:

docker-compose up -d
This command will start a container with PostgreSQL and another container with the Go application.

Accessing Swagger
Once the application is running, whether directly or through Docker, you can access the Swagger documentation at the following URL:

http://localhost:50020/swagger/index.html#/

Project Configuration
Environment Variables
The project uses the following environment variables that need to be properly configured for the database and server execution:

POSTGRES_DB_HOST=localhost
POSTGRES_DB_PORT=5433
POSTGRES_DB_USER=postgres
POSTGRES_DB_PASSWORD=postgres
POSTGRES_DB_NAME=postgres
POSTGRES_DB_SSLMODE=disable
GORM_MODE=on
GIN_MODE=debug
SERVER_PORT="localhost:50020"
These variables are already configured in the .env file, which is included in the container when running with Docker.

Documentation
You can review the full API documentation, including all available endpoints and their details, directly in the Swagger interface mentioned above.

Contact
If you have any questions or suggestions, feel free to contact me.# leal-technical-test
Technical test
