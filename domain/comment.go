package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserID  *uint  `json:"user_id"`
	PhotoID *uint  `json:"photo_id"`
	Message string `json:"message" gorm:"not null"`
	User    *User  `gorm:"foreignkey:UserID"`
	Photo   *Photo `gorm:"foreignkey:PhotoID"`
}

func (comment Comment) Validate() error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.Message, validation.Required),
	)
}
