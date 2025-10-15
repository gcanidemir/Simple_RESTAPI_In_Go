package handler

import (
	"Simple_RESTAPI_In_Go/dto"
	"Simple_RESTAPI_In_Go/entity"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginUser(c *gin.Context) {
	var loginDTO dto.LoginDTO

	if err := c.BindJSON(&loginDTO); err != nil {
		return
	}

	for _, user := range entity.Users {
		if loginDTO.Username == user.Username && loginDTO.Password == user.Password {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": jwt.NewNumericDate(time.Now().Add(5 * time.Hour)), // Token expires in 5 hours
			})

			jwtSecret := []byte(os.Getenv("JWT_SECRET"))
			tokenString, err := token.SignedString(jwtSecret)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
			return
		}
	}
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
}

func RegisterUser(c *gin.Context) {
	var newUser entity.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	entity.Users = append(entity.Users, newUser)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": jwt.NewNumericDate(time.Now().Add(5 * time.Hour)), // Token expires in 5 hours
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("jwt-secret")))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"token": tokenString})
}
