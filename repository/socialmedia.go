package repository

import (
	"MyGram/database"
	"MyGram/domain"
)

type SocialMediaRepository struct {
}

func (s *SocialMediaRepository) Create(socialMedia domain.SocialMedia) (domain.RespCreateSocialMedia, error) {
	result := domain.RespCreateSocialMedia{}
	err := database.Database.DB.Create(&socialMedia).Error
	if err != nil {
		return result, err
	}
	result.Id = int(socialMedia.ID)
	result.Name = socialMedia.Name
	result.SocialMediaUrl = socialMedia.SocialMediaUrl
	result.UserId = int(*socialMedia.UserId)
	result.CreatedAt = socialMedia.CreatedAt.String()
	return result, err
}

func (s *SocialMediaRepository) Get() (domain.RespGetSocialMedia, error) {
	result := domain.RespGetSocialMedia{}
	err := database.Database.DB.Table("social_medias").Where("deleted_at IS NULL").Scan(&result.SocialMedias).Error
	if err != nil {
		return result, err
	}
	for r := range result.SocialMedias {
		err = database.Database.DB.Table("users").Where("id = ?", result.SocialMedias[r].UserId).Where("deleted_at IS NULL").Scan(&result.SocialMedias[r].User).Error
		if err != nil {
			return result, err
		}
	}
	return result, err
}

func (s *SocialMediaRepository) Update(socialMedia domain.SocialMedia) (domain.RespUpdateSocialMedia, error) {
	result := domain.RespUpdateSocialMedia{}
	err := database.Database.DB.Table("social_medias").Where("id = ?", socialMedia.ID).Where("deleted_at IS NULL").Updates(&socialMedia).Error
	if err != nil {
		return result, err
	}
	err = database.Database.DB.Table("social_medias").Where("id = ?", socialMedia.ID).Where("deleted_at IS NULL").First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}

func (s *SocialMediaRepository) Delete(id int) (bool, error) {
	socialMedia := domain.SocialMedia{}
	err := database.Database.DB.Where("id = ?", id).Where("deleted_at IS NULL").First(&socialMedia).Error
	if err != nil {
		return false, err
	}
	err = database.Database.DB.Delete(&socialMedia).Error
	if err != nil {
		return false, err
	}
	return true, err
}
