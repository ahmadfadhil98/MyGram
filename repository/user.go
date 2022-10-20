package repository

import (
	"MyGram/database"
	"MyGram/domain"
	"errors"
	"fmt"
)

type UserRepository struct {
}

func (u *UserRepository) Register(user domain.User) (domain.RespRegister, error) {
	result := domain.RespRegister{}
	err := database.Database.DB.Create(&user).Error
	//err := idb.DB.Create(&user).Error
	//check if error is 23505

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			return result, errors.New("username already exists")
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return result, errors.New("email already exists")
		}
		return result, err
	}
	result.Id = int(user.ID)
	result.Email = user.Email
	result.Username = *user.Username
	result.Age = *user.Age

	return result, err

}

func (u *UserRepository) Login(user domain.User) (bool, domain.User, error) {
	var result domain.User
	err := database.Database.DB.Table("users").Where("email = ?", user.Email).Where("deleted_at IS NULL").First(&result).Error
	if err != nil {
		return false, result, err
	}
	return true, result, err
}

func (u *UserRepository) Update(user domain.User) (domain.RespUpdateUser, error) {
	result := domain.RespUpdateUser{}
	err := database.Database.DB.Table("users").Where("id = ?", user.ID).Where("deleted_at IS NULL").
		Updates(domain.User{
			Username: user.Username,
			Email:    user.Email,
		}).Error
	if err != nil {
		return result, err
	}
	err = database.Database.DB.Table("users").Where("id = ?", user.ID).Where("deleted_at IS NULL").First(&result).Error
	fmt.Println(result)
	//result.Id = int(user.ID)
	//result.Email = user.Email
	//result.Username = *user.Username
	//result.Age = *user.Age
	//result.UpdatedAt = user.UpdatedAt.String()
	return result, err
}

func (u *UserRepository) Delete(id int64) (bool, error) {
	user := domain.User{}
	err := database.Database.DB.Where("id = ?", id).Where("deleted_at IS NULL").First(&user).Error
	if err != nil {
		return false, err
	}
	err = database.Database.DB.Delete(&user).Error
	if err != nil {
		return false, err
	}
	return true, err
}
