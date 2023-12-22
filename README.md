# json api server

This is a simple JSON API server. The server can perform basic operations such as creating an account, viewing accounts, and deleting an account. It also includes a simple authorization system based on JSON Web Tokens (JWT).

## Getting Started

To run this application, install the Go programming language and the required dependencies.

### Prerequisites

- Go: [https://golang.org/dl/](https://golang.org/dl/)
- Install the required packages by running the following command in the terminal within the project directory:

  ```bash
  go get -d -v ./...
Running
To start the application, run the following command in the terminal within the project directory:

bash
Copy code
go run main.go
The application will run by default at http://localhost:8080.

API Endpoints
View Accounts
URL: /account
Method: GET
Description: Displays all accounts.
Create Account
URL: /account
Method: POST
Description: Creates a new account and returns a JSON Web Token (JWT).
View Account
URL: /account/{id}
Method: GET
Description: Displays a specific account based on the ID.
Delete Account
URL: /account/{id}
Method: DELETE
Description: Deletes a specific account based on the ID.
JWT Authorization
JWT is used to access API endpoints. Each request must include a valid JWT through the x-jwt-token header.

Environment Variables
JWT_SECRET: Secret key to be used for JWT.
To set example environment variables, run the following command in the terminal:

bash
Copy code
export JWT_SECRET="omerfdev0101"
