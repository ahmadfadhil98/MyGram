package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username    *string        `json:"username" gorm:"type:varchar(100);not null;unique"`
	Email       string         `json:"email" gorm:"type:varchar(100);not null;unique"`
	Password    *string        `json:"password"`
	Age         *int           `json:"age"`
	Photos      *[]Photo       `gorm:"foreignkey:UserId"`
	SocialMedia *[]SocialMedia `gorm:"foreignkey:UserId"`
	Comments    *[]Comment     `gorm:"foreignkey:UserId"`
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 225)),
		validation.Field(&user.Age, validation.Required, validation.Min(8)),
	)
}
