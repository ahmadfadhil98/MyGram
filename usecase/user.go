package usecase

import (
	"MyGram/domain"
	"MyGram/helper"
	"MyGram/repository"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserUsecase struct {
	userRepo repository.UserRepository
	jwt      helper.Jwt
}

func (u *UserUsecase) Register(c *gin.Context) {
	var user domain.User
	var response domain.Response
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = user.Validate()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	*user.Password = string(hashedPassword)
	result, err := u.userRepo.Register(user)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusCreated
	response.Data = result
	c.JSON(http.StatusCreated, response)
}

func (u *UserUsecase) Login(c *gin.Context) {
	var user domain.User
	var result domain.RespLogin
	var response domain.Response
	err := godotenv.Load("../MyGram/database/.env")
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = c.ShouldBindJSON(&user)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	exist, resp, err := u.userRepo.Login(user)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if exist {
		err = bcrypt.CompareHashAndPassword([]byte(*resp.Password), []byte(*user.Password))
		if err != nil {
			response.Status = http.StatusUnauthorized
			response.Data = gin.H{"error": "Password is wrong"}
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		token, err := u.jwt.GetToken(resp.ID)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Data = gin.H{"error": err.Error()}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		result.Token = token
	}
	response.Status = http.StatusOK
	response.Data = result
	c.JSON(http.StatusOK, response)
}

func (u *UserUsecase) Update(c *gin.Context) {
	var user domain.User
	var response domain.Response
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	id := c.Param("userId")
	u64, _ := strconv.ParseUint(id, 10, 32)
	user.ID = uint(u64)
	sesi := sessions.Default(c)
	akses, data, err := u.jwt.CheckToken(sesi)
	if err != nil {
		response.Status = http.StatusInternalServerError
		if err.Error() == "record not found" {
			response.Data = gin.H{"error": "You must login first"}
		} else {
			response.Data = gin.H{"error": err.Error()}
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if akses {
		user.ID = uint(data.Client_Id)
		result, err := u.userRepo.Update(user)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Data = gin.H{"error": err.Error()}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response.Status = http.StatusOK
		response.Data = result
		c.JSON(http.StatusOK, response)
	}

}

func (u *UserUsecase) Delete(c *gin.Context) {
	var response domain.Response
	sesi := sessions.Default(c)
	akses, dataAkses, err := u.jwt.CheckToken(sesi)
	if err != nil {
		response.Status = http.StatusInternalServerError
		if err.Error() == "record not found" {
			response.Data = gin.H{"error": "You must login first"}
		} else {
			response.Data = gin.H{"error": err.Error()}
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if akses {
		success, err := u.userRepo.Delete(dataAkses.Client_Id)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Data = gin.H{"error": err.Error()}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		if success {
			response.Status = http.StatusOK
			response.Data = gin.H{"message": "Your account has been successfully deleted"}
			c.JSON(http.StatusOK, response)
			return
		}
	}

}
