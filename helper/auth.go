package helper

import (
	"MyGram/domain"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type Auth struct {
}

func (a *Auth) Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")

	err := godotenv.Load("../MyGram/database/.env")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if token != nil && err == nil {
		session := sessions.Default(c)
		session.Set("claims", token.Claims)
		err = session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	} else {
		result := domain.Response{}
		result.Status = http.StatusUnauthorized
		result.Data = gin.H{"error": "Unauthorized"}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
		return
	}
}
