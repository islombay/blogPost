package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/islombay/blogPost/internal/service"
	"github.com/islombay/blogPost/pkg/httpserver/handlers/post_handler"
)

type Handler struct {
	Post post_handler.PostHandlerInterface
}

func InitRoutes(s *service.BlogPostService) *gin.Engine {
	r := gin.Default()

	hnd := Handler{
		Post: post_handler.NewPostHandler(s.Post),
	}

	apiGroup := r.Group("/api")
	{
		post := apiGroup.Group("/post")
		{
			post.POST("/new", hnd.Post.HandlerCreate)

			post.GET("/all", hnd.Post.HandlerGetAll)

			post.DELETE("/:id", hnd.Post.HandlerDelete)
		}
	}

	return r
}
