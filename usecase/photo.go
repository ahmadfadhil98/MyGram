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

type PhotoUsecase struct {
	photoRepo repository.PhotoRepository
	jwt       helper.Jwt
}

func (p *PhotoUsecase) Create(c *gin.Context) {
	var photo domain.Photo
	var response domain.Response
	err := c.ShouldBindJSON(&photo)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = photo.Validate()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	sesi := sessions.Default(c)
	akses, data, err := p.jwt.CheckToken(sesi)
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
		photo.UserId = new(uint)
		*photo.UserId = uint(data.Client_Id)
		result, err := p.photoRepo.Create(photo)
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

func (p *PhotoUsecase) Get(c *gin.Context) {
	var response domain.Response
	sesi := sessions.Default(c)
	akses, _, err := p.jwt.CheckToken(sesi)
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
		result, err := p.photoRepo.Get()
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

func (p *PhotoUsecase) Update(c *gin.Context) {
	var photo domain.Photo
	var response domain.Response
	err := c.ShouldBindJSON(&photo)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	id := c.Param("photoId")
	u64, _ := strconv.ParseUint(id, 10, 32)
	photo.ID = uint(u64)
	sesi := sessions.Default(c)
	akses, data, err := p.jwt.CheckToken(sesi)
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
		photo.UserId = new(uint)
		*photo.UserId = uint(data.Client_Id)
		result, err := p.photoRepo.Update(photo)
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

func (p *PhotoUsecase) Delete(c *gin.Context) {
	var response domain.Response
	id := c.Param("photoId")
	photoId, err := strconv.Atoi(id)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	sesi := sessions.Default(c)
	akses, _, err := p.jwt.CheckToken(sesi)
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
		success, err := p.photoRepo.Delete(photoId)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Data = gin.H{"error": err.Error()}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		if success {
			response.Status = http.StatusOK
			response.Data = gin.H{"message": "Your photo has been successfully deleted"}
			c.JSON(http.StatusOK, response)
			return
		}
	}
}
