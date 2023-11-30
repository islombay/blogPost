package post_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"github.com/islombay/blogPost/pkg/utils/objects"
	"log/slog"
	"net/http"
	"time"
)

type HandlerGetAllResPostBody struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username,omitempty"`
}

type HandlerGetAllResBody struct {
	Posts []HandlerGetAllResPostBody `json:"posts"`
}

func (h *PostHandler) HandlerGetAll(c *gin.Context) {
	all, err := h.service.GetAll()
	if err != nil {
		switch err.Error() {
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
	}
	res := []HandlerGetAllResPostBody{}

	//for _, e := range all {
	//	data, err := json.Marshal(e)
	//	if err != nil {
	//		slog.Error("could not marshal in handler get all posts", sl.Err(err))
	//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
	//		return
	//	}
	//	n := HandlerGetAllResPostBody{}
	//	err = json.Unmarshal(data, &n)
	//	if err != nil {
	//		slog.Error("could not unmarshal in handler get all posts", sl.Err(err))
	//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
	//		return
	//	}
	//	res = append(res, n)
	//}

	for _, e := range all {
		n := HandlerGetAllResPostBody{}
		err := objects.ReObject(e, &n)
		if err != nil {
			slog.Error("could not (un)marshal in handler get all posts", sl.Err(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
		res = append(res, n)
	}
	c.JSON(http.StatusOK, HandlerGetAllResBody{Posts: res})
}
