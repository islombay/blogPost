package post_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/islombay/blogPost/internal/service/post"
)

type PostHandler struct {
	service post.PostServiceInterface
}

type PostHandlerInterface interface {
	HandlerCreate(c *gin.Context)
	HandlerGetAll(c *gin.Context)
	HandlerDelete(c *gin.Context)
}

func NewPostHandler(postService post.PostServiceInterface) PostHandlerInterface {
	return &PostHandler{service: postService}
}
