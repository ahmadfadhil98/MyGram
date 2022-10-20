package repository

import (
	"MyGram/database"
	"MyGram/domain"
)

type CommentRepository struct {
}

func (c *CommentRepository) Create(comment domain.Comment) (domain.RespCreateComment, error) {
	result := domain.RespCreateComment{}
	err := database.Database.DB.Create(&comment).Error
	if err != nil {
		return result, err
	}
	result.Id = int(comment.ID)
	result.Message = comment.Message
	result.PhotoId = int(*comment.PhotoID)
	result.UserId = int(*comment.UserID)
	result.CreatedAt = comment.CreatedAt.String()
	return result, err
}

func (c *CommentRepository) Get() ([]domain.RespGetComment, error) {
	result := []domain.RespGetComment{}
	err := database.Database.DB.Table("comments").Where("deleted_at IS NULL").Scan(&result).Error
	if err != nil {
		return result, err
	}
	for r := range result {
		err = database.Database.DB.Table("users").Where("id = ?", result[r].UserId).Scan(&result[r].User).Error
		if err != nil {
			return result, err
		}
		err = database.Database.DB.Table("photos").Where("id = ?", result[r].PhotoId).Scan(&result[r].Photo).Error
		if err != nil {
			return result, err
		}
	}
	return result, err
}

func (c *CommentRepository) Update(comment domain.Comment) (domain.RespUpdateComment, error) {
	result := domain.RespUpdateComment{}
	err := database.Database.DB.Table("comments").Where("id = ?", comment.ID).Where("deleted_at IS NULL").Updates(&comment).Error
	if err != nil {
		return result, err
	}
	err = database.Database.DB.Table("comments").Where("id = ?", comment.ID).Where("deleted_at IS NULL").First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}

func (c *CommentRepository) Delete(id int) (bool, error) {
	comment := domain.Comment{}
	err := database.Database.DB.Where("id = ?", id).Where("deleted_at IS NULL").First(&comment).Error
	if err != nil {
		return false, err
	}
	err = database.Database.DB.Delete(&comment).Error
	if err != nil {
		return false, err
	}
	return true, err
}
