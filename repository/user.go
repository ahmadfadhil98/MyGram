package repository

import (
	"MyGram/domain"
	"MyGram/infrastructure"
	"errors"
)

type UserRepository struct {
}

func (idb *UserRepository) Register(user domain.User) (domain.RespRegister, error) {
	result := domain.RespRegister{}
	err := infrastructure.Database.DB.Create(&user).Error
	//err := idb.DB.Create(&user).Error
	//check if error is 23505

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			return result, errors.New("username already exists")
		}
	}
	result.Id = user.ID
	result.Email = user.Email
	result.Username = user.Username
	result.Age = user.Age

	return result, err

}

func (idb *UserRepository) Login(user domain.ReqLogin) (bool, domain.ReqLogin, error) {
	var result domain.ReqLogin
	err := infrastructure.Database.DB.Where("email = ?", user.Email).First(&result).Error
	if err != nil {
		return false, result, err
	}
	return true, result, err
}
