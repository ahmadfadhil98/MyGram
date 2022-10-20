package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"type:varchar(100);not null"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null"`
	UserId         *uint  `json:"user_id"`
	User           *User  `gorm:"foreignkey:UserId"`
}

func (socialMedia SocialMedia) Validate() error {
	return validation.ValidateStruct(&socialMedia,
		validation.Field(&socialMedia.Name, validation.Required),
		validation.Field(&socialMedia.SocialMediaUrl, validation.Required),
	)
}
