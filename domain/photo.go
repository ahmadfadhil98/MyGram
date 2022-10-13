package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `form:"title"`
	Caption  string `form:"caption"`
	PhotoUrl string `form:"photo_url"`
	UserId   uint   `form:"user_id"`
	User     User   `gorm:"foreignkey:UserId"`
}

func (photo Photo) Validate() error {
	return validation.ValidateStruct(&photo,
		validation.Field(&photo.Title, validation.Required),
		validation.Field(&photo.PhotoUrl, validation.Required),
	)
}
