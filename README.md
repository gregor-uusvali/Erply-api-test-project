# Erply API Test Project

This project is a simple Golang-based Erply API endpoint/middleware that allows reading and writing customer data from the Erply API. Users can log in with their Erply credentials (clientCode, username, password) to add new customers, view the customers table, and delete customers.

## Requirements

To run this project, make sure you have the following:

- Go programming language (version 1.15 or above) installed
- Docker installed (if you plan to use Docker for containerization)
- Erply API credentials (client code, username, and password)

## Installation

1. Clone the repository:
   `git clone https://github.com/gregor-uusvali/Erply-api-test-project`
   `cd Erply-api-test-project`

## Usage

Run the application:

Run in terminal
`go run .`,

Run the executable
`./Erply-api-test-project `

Docker
`./run_docker.sh`

The application will start and listen on http://localhost:8080.

Access the API documentation:

Open your web browser and navigate to http://localhost:8080/swagger/.

Unfortunately, I couldn't get the Swagger UI to display my API documentation, but the functions have been annotated.

## Structure

The project structure:

- main.go: The entry point of the application that sets up the server and routes.
- driver: Contains the database connection and initialization logic.
- handlers: Handles the API endpoints and performs business logic.
- api: Handles the communication with the Erply API.
- repository: Provides the database repository interfaces.
- repository/dbrepo: Implements the database repository using SQLite.

## Tests

Unit tests for database operations are included in `/repository/dbrepo/customerrepo_test.go`

### Author

[Gregor Uusväli](https://github.com/gregor-uusvali)
