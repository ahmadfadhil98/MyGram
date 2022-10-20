package repository

import (
	"MyGram/database"
	"MyGram/domain"
)

type PhotoRepository struct {
}

func (p *PhotoRepository) Create(photo domain.Photo) (domain.RespCreatePhoto, error) {
	result := domain.RespCreatePhoto{}
	err := database.Database.DB.Create(&photo).Error
	if err != nil {
		return result, err
	}
	result.Id = int(photo.ID)
	result.Title = photo.Title
	result.Caption = *photo.Caption
	result.PhotoUrl = photo.PhotoUrl
	result.CreatedAt = photo.CreatedAt.String()
	return result, err
}

func (p *PhotoRepository) Get() ([]domain.RespGetPhoto, error) {
	result := []domain.RespGetPhoto{}
	err := database.Database.DB.Table("photos").Where("deleted_at IS NULL").Scan(&result).Error
	if err != nil {
		return result, err
	}
	for r := range result {
		err = database.Database.DB.Table("users").Where("id = ?", result[r].UserId).Where("deleted_at IS NULL").Scan(&result[r].User).Error
		if err != nil {
			return result, err
		}
	}
	return result, err
}

func (p *PhotoRepository) GetById(id int) ([]domain.RespGetPhoto, error) {
	result := []domain.RespGetPhoto{}
	err := database.Database.DB.Table("photos").Where("id = ? AND deleted_at IS NULL", id).Scan(&result).Error
	if err != nil {
		return result, err
	}
	for r := range result {
		err = database.Database.DB.Table("users").Where("id = ?", result[r].UserId).Where("deleted_at IS NULL").Scan(&result[r].User).Error
		if err != nil {
			return result, err
		}
	}
	return result, err
}

func (p *PhotoRepository) Update(photo domain.Photo) (domain.RespUpdatePhoto, error) {
	result := domain.RespUpdatePhoto{}
	err := database.Database.DB.Table("photos").Where("id = ?", photo.ID).Where("deleted_at IS NULL").Updates(&photo).Error
	if err != nil {
		return result, err
	}
	err = database.Database.DB.Table("photos").Where("id = ?", photo.ID).Where("deleted_at IS NULL").First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}

func (p *PhotoRepository) Delete(id int) (bool, error) {
	photo := domain.Photo{}
	err := database.Database.DB.Where("id = ?", id).Where("deleted_at IS NULL").First(&photo).Error
	if err != nil {
		return false, err
	}
	err = database.Database.DB.Delete(&photo).Error
	if err != nil {
		return false, err
	}
	return true, err
}
