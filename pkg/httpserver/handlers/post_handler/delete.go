package post_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *PostHandler) HandlerDelete(c *gin.Context) {
	post_id := c.Param("id")
	intValue, err := strconv.ParseInt(post_id, 10, 64)

	if err != nil {
		slog.Error("could not convert string to int64 in post delete", sl.Err(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}
	err = h.service.Delete(intValue)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
