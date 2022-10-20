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

		users := router.Group("/users")
		users.Use(r.Auth.Auth)
		{
			users.PUT("/", r.UserUse.Update)
			users.DELETE("/", r.UserUse.Delete)
		}

		photos := router.Group("/photos")
		photos.Use(r.Auth.Auth)
		{
			photos.POST("/", r.PhotoUse.Create)
			photos.GET("/", r.PhotoUse.Get)
			photos.PUT("/:photoId", r.PhotoUse.Update)
			photos.DELETE("/:photoId", r.PhotoUse.Delete)
		}
		comments := router.Group("/comments")
		comments.Use(r.Auth.Auth)
		{
			comments.POST("/", r.CommentUse.Create)
			comments.GET("/", r.CommentUse.Get)
			comments.PUT("/:commentId", r.CommentUse.Update)
			comments.DELETE("/:commentId", r.CommentUse.Delete)
		}
		socmeds := router.Group("/socialmedias")
		socmeds.Use(r.Auth.Auth)
		{
			socmeds.POST("/", r.SocmedUse.Create)
			socmeds.GET("/", r.SocmedUse.Get)
			socmeds.PUT("/:socialMediaId", r.SocmedUse.Update)
			socmeds.DELETE("/:socialMediaId", r.SocmedUse.Delete)
		}

	}

	err := router.Run(":8001")
	if err != nil {
		return nil, err
	}
	return router, nil
}
