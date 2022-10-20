package helper

import (
	"MyGram/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/mitchellh/mapstructure"
	"os"
	"time"
)

type Jwt struct {
	user repository.UserRepository
}

type TokenClaim struct {
	Client_Id int64
}

func (j *Jwt) GetToken(id uint) (token string, err error) {
	sign := jwt.New(jwt.SigningMethodHS256)
	claims := sign.Claims.(jwt.MapClaims)
	claims["client_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token, err = sign.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token, err
}

func (j *Jwt) CheckToken(s sessions.Session) (akses bool, data TokenClaim, err error) {
	tokenClaim := TokenClaim{}

	if s == nil {
		return false, tokenClaim, fmt.Errorf("Session is empty")
	}

	claims := s.Get("claims")
	if claims == nil {
		return false, tokenClaim, fmt.Errorf("Claims is empty")
	}
	err = mapstructure.Decode(claims, &tokenClaim)
	if err != nil {
		return false, tokenClaim, err
	}
	success, err := j.user.GetById(tokenClaim.Client_Id)
	if err != nil {
		return false, tokenClaim, err
	}
	if !success {
		return false, tokenClaim, fmt.Errorf("User not found")
	}
	return true, tokenClaim, nil
}
