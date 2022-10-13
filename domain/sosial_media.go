package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `form:"name" `
	SocialMediaUrl string `form:"social_media_url"`
	UserId         uint   `form:"user_id"`
	User           User   `gorm:"foreignkey:UserId"`
}

func (socialMedia SocialMedia) Validate() error {
	return validation.ValidateStruct(&socialMedia,
		validation.Field(&socialMedia.Name, validation.Required),
		validation.Field(&socialMedia.SocialMediaUrl, validation.Required),
	)
}
