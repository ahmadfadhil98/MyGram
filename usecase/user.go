package usecase

import (
	"MyGram/domain"
	"MyGram/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func (userusecase *UserUsecase) Register(c *gin.Context) {
	var user domain.User
	c.Bind(&user)
	err := user.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	result, err := userusecase.userRepo.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (userusecase *UserUsecase) Login(c *gin.Context) {
	var user domain.ReqLogin
	var result domain.RespLogin
	err := godotenv.Load("../MyGram/app/.env")
	c.Bind(&user)
	exist, resp, err := userusecase.userRepo.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exist {
		err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(user.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		sign := jwt.New(jwt.SigningMethodHS256)
		claims := sign.Claims.(jwt.MapClaims)
		claims["email"] = user.Email
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		token, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		result.Token = token
	}
	c.JSON(http.StatusOK, result)
}
