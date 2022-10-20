package usecase

import (
	"MyGram/domain"
	"MyGram/helper"
	"MyGram/repository"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentUsecase struct {
	commentRepo repository.CommentRepository
	photoRepo   repository.PhotoRepository
	jwt         helper.Jwt
}

func (cu *CommentUsecase) Create(c *gin.Context) {
	var comment domain.Comment
	var response domain.Response
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = comment.Validate()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	photoId := int(*comment.PhotoID)
	fmt.Println(photoId)
	photo, err := cu.photoRepo.GetById(photoId)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if len(photo) == 0 {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": "Photo ID not found"}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	sesi := sessions.Default(c)
	akses, data, err := cu.jwt.CheckToken(sesi)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if akses {
		comment.UserID = new(uint)
		*comment.UserID = uint(data.Client_Id)
		result, err := cu.commentRepo.Create(comment)
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

func (cu *CommentUsecase) Get(c *gin.Context) {
	var response domain.Response
	sesi := sessions.Default(c)
	akses, _, err := cu.jwt.CheckToken(sesi)
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
		result, err := cu.commentRepo.Get()
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

func (cu *CommentUsecase) Update(c *gin.Context) {
	var comment domain.Comment
	var response domain.Response
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	id := c.Param("commentId")
	u64, _ := strconv.ParseUint(id, 10, 32)
	comment.ID = uint(u64)
	sesi := sessions.Default(c)
	akses, data, err := cu.jwt.CheckToken(sesi)
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
		comment.UserID = new(uint)
		*comment.UserID = uint(data.Client_Id)
		fmt.Println(comment)
		result, err := cu.commentRepo.Update(comment)
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

func (cu *CommentUsecase) Delete(c *gin.Context) {
	var response domain.Response
	id := c.Param("commentId")
	commentId, err := strconv.Atoi(id)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	sesi := sessions.Default(c)
	akses, _, err := cu.jwt.CheckToken(sesi)
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
		success, err := cu.commentRepo.Delete(commentId)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Data = gin.H{"error": err.Error()}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		if success {
			response.Status = http.StatusOK
			response.Data = gin.H{"message": "Your comment has been successfully deleted"}
			c.JSON(http.StatusOK, response)
			return
		}
	}

}
