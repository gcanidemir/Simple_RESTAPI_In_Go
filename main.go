package main

import (
	"Simple_RESTAPI_In_Go/handler"
	"Simple_RESTAPI_In_Go/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	loadEnv()

	router := gin.Default()

	router.GET("/books", middleware.AuthMiddleware(), handler.GetBooks)
	router.GET("/books/:id", middleware.AuthMiddleware(), handler.GetBookByID)
	router.POST("/books", middleware.AuthMiddleware(), handler.PostBook)
	router.PUT("/books/:id", middleware.AuthMiddleware(), handler.UpdateBookByID)
	router.DELETE("/books/:id", middleware.AuthMiddleware(), handler.DeleteBookByID)

	router.POST("/users/login", handler.LoginUser)
	router.POST("/users/register", handler.RegisterUser)

	router.Run("localhost:8080")
}
