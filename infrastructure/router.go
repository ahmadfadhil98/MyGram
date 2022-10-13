package infrastructure

import (
	"MyGram/usecase"
	"github.com/gin-gonic/gin"
)

var Route Router

type Router struct {
	UserUse *usecase.UserUsecase
}

func (r *Router) RouterInit() (*gin.Engine, error) {

	router := gin.Default()
	router.POST("/register", r.UserUse.Register)

	err := router.Run(":8080")
	if err != nil {
		return nil, err
	}
	return router, nil
}
