package infrastructure

import (
	"MyGram/helper"
	"MyGram/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

var Route Router

type Router struct {
	UserUse    usecase.UserUsecase
	PhotoUse   usecase.PhotoUsecase
	CommentUse usecase.CommentUsecase
	SocmedUse  usecase.SocialMediaUsecase
	Auth       helper.Auth
}

func (r *Router) RouterInit() (*gin.Engine, error) {

	router := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("JWT_SECRET")))
	router.Use(sessions.Sessions("backend", store))
	{
		router.POST("/register", r.UserUse.Register)
		router.POST("/login", r.UserUse.Login)
		router.Use(r.Auth.Auth)
		{
			router.PUT("/users", r.UserUse.Update)
			router.DELETE("/users", r.UserUse.Delete)
		}

		router.Use(r.Auth.Auth)
		{
			router.POST("/photos", r.PhotoUse.Create)
			router.GET("/photos", r.PhotoUse.Get)
			router.PUT("/photos/:photoId", r.PhotoUse.Update)
			router.DELETE("/photos/:photoId", r.PhotoUse.Delete)
		}
		router.Use(r.Auth.Auth)
		{
			router.POST("/comments", r.CommentUse.Create)
			router.GET("/comments", r.CommentUse.Get)
			router.PUT("/comments/:commentId", r.CommentUse.Update)
			router.DELETE("/comments/:commentId", r.CommentUse.Delete)
		}
		router.Use(r.Auth.Auth)
		{
			router.POST("/socmed", r.SocmedUse.Create)
			router.GET("/socmed", r.SocmedUse.Get)
			router.PUT("/socmed/:socialMediaId", r.SocmedUse.Update)
			router.DELETE("/socmed/:socialMediaId", r.SocmedUse.Delete)
		}

	}

	err := router.Run(":8001")
	if err != nil {
		return nil, err
	}
	return router, nil
}
