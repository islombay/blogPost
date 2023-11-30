package post_handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/islombay/blogPost/internal/service/post"
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"log/slog"
	"net/http"
	"time"
)

type HandlerCreateBody struct {
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	Username  string    `json:"username"`
}

type HandlerCreateResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username,omitempty"`
}

func (h *PostHandler) HandlerCreate(c *gin.Context) {
	var body HandlerCreateBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "Post Body is not correct"},
		)
		return
	}

	postModel, err := h.service.Create(post.CreateModel{
		Title:     body.Title,
		Content:   body.Content,
		CreatedAt: body.CreatedAt,
		Username:  body.Username,
	})

	if err != nil {
		switch err.Error() {
		case post.InvalidModel:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Post Body is not correct"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
	}

	data, err := json.Marshal(postModel)
	if err != nil {
		slog.Error("could not marshal in create handler", sl.Err(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	var v HandlerCreateResponse
	if err := json.Unmarshal(data, &v); err != nil {
		slog.Error("could not unmarshal in create handler", sl.Err(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, v)
}
