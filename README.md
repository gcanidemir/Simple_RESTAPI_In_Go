# Simple_RESTAPI_In_Go

A simple REST API built with Go and the Gin framework. This project provides endpoints for managing books and users, with JWT-based authentication.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.15+ recommended)
- A `.env` file in the root directory with your environment variables.

### Installation

1.  Clone the repository:
    ```sh
    git clone https://github.com/gcanidemir/Simple_RESTAPI_In_Go.git
    ```
2.  Navigate to the project directory:
    ```sh
    cd Simple_RESTAPI_In_Go
    ```
3.  Create a `.env` file in the root of the project and add the following content:
    ```
    JWT_SECRET=your_secret_key
    ```
4.  Install the dependencies:
    ```sh
    go get .
    ```

### Running the Application

To run the application, execute the following command from the root directory:

```sh
go run main.go
```

The server will start on `localhost:8080`.

## API Endpoints

### User Endpoints

-   `POST /users/register`: Register a new user.
-   `POST /users/login`: Login a user and receive a JWT token.

### Book Endpoints

All book endpoints require a valid JWT token in the `Authorization` header.

-   `GET /books`: Get a list of all books.
-   `GET /books/:id`: Get a single book by its ID.
-   `POST /books`: Create a new book.
-   `PUT /books/:id`: Update a book by its ID.
-   `DELETE /books/:id`: Delete a book by its ID.

## Dependencies

-   [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin): A popular web framework for Go.
-   [github.com/joho/godotenv](https://github.com/joho/godotenv): A Go dotenv library
-   [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt): A Go implementation of JSON Web Tokens.
