package routes

import (
	"first-go-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	postRoutes := router.Group("api/v1/posts")
	{
		postRoutes.POST("/", controllers.CreatePost)
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.GET("/:id", controllers.GetPost)
		postRoutes.PUT("/:id", controllers.UpdatePost)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}
}
