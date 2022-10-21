package usecase

import (
	"MyGram/domain"
	"MyGram/helper"
	"MyGram/repository"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SocialMediaUsecase struct {
	socialmediaRepo repository.SocialMediaRepository
	jwt             helper.Jwt
}

func (s *SocialMediaUsecase) Create(c *gin.Context) {
	var socialmedia domain.SocialMedia
	var response domain.Response
	err := c.ShouldBindJSON(&socialmedia)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = socialmedia.Validate()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	sesi := sessions.Default(c)
	akses, data, err := s.jwt.CheckToken(sesi)
	if err != nil {
		response.Status = http.StatusBadRequest
		if err.Error() == "record not found" {
			response.Data = gin.H{"error": "You must login first"}
		} else {
			response.Data = gin.H{"error": err.Error()}
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if akses {
		socialmedia.UserId = new(uint)
		*socialmedia.UserId = uint(data.Client_Id)
		result, err := s.socialmediaRepo.Create(socialmedia)
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
}

func (s *SocialMediaUsecase) Get(c *gin.Context) {
	var response domain.Response
	sesi := sessions.Default(c)
	akses, _, err := s.jwt.CheckToken(sesi)
	if err != nil {
		response.Status = http.StatusBadRequest
		if err.Error() == "record not found" {
			response.Data = gin.H{"error": "You must login first"}
		} else {
			response.Data = gin.H{"error": err.Error()}
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if akses {
		result, err := s.socialmediaRepo.Get()
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

func (s *SocialMediaUsecase) Update(c *gin.Context) {
	var socialmedia domain.SocialMedia
	var response domain.Response
	err := c.ShouldBindJSON(&socialmedia)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = socialmedia.Validate()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	id := c.Param("socialMediaId")
	u64, _ := strconv.ParseUint(id, 10, 32)
	socialmedia.ID = uint(u64)
	sesi := sessions.Default(c)
	akses, data, err := s.jwt.CheckToken(sesi)
	if err != nil {
		response.Status = http.StatusBadRequest
		if err.Error() == "record not found" {
			response.Data = gin.H{"error": "You must login first"}
		} else {
			response.Data = gin.H{"error": err.Error()}
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if akses {
		socialmedia.UserId = new(uint)
		*socialmedia.UserId = uint(data.Client_Id)
		result, err := s.socialmediaRepo.Update(socialmedia)
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

func (s *SocialMediaUsecase) Delete(c *gin.Context) {
	var response domain.Response
	id := c.Param("socialMediaId")
	socialMediaId, err := strconv.Atoi(id)
	sesi := sessions.Default(c)
	akses, _, err := s.jwt.CheckToken(sesi)
	if err != nil {
		response.Status = http.StatusBadRequest
		if err.Error() == "record not found" {
			response.Data = gin.H{"error": "You must login first"}
		} else {
			response.Data = gin.H{"error": err.Error()}
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if akses {
		success, err := s.socialmediaRepo.Delete(socialMediaId)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Data = gin.H{"error": err.Error()}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		if success {
			response.Status = http.StatusOK
			response.Data = gin.H{"message": "Your social media has been successfully deleted"}
			c.JSON(http.StatusOK, response)
		}
	}
}
