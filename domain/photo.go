package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Photo struct {
	gorm.Model
	Title    string  `json:"title" gorm:"type:varchar(100);not null"`
	Caption  *string `json:"caption" gorm:"type:varchar(100)"`
	PhotoUrl string  `json:"photo_url" gorm:"not null"`
	UserId   *uint   `json:"user_id"`
	User     *User   `gorm:"foreignkey:UserId"`
}

func (photo Photo) Validate() error {
	return validation.ValidateStruct(&photo,
		validation.Field(&photo.Title, validation.Required),
		validation.Field(&photo.PhotoUrl, validation.Required),
	)
}
