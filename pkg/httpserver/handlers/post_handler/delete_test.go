package post_handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/islombay/blogPost/internal/service"
	mock_post "github.com/islombay/blogPost/internal/service/post/mocks"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestPostHandler_HandlerDelete(t *testing.T) {
	type mockBehaviour func(s *mock_post.MockPostServiceInterface, id int64)

	testTable := []struct {
		name           string
		postID         string
		postIDInt      int64
		mockBehaviour  mockBehaviour
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "OK",
			postID:    "2",
			postIDInt: 2,
			mockBehaviour: func(s *mock_post.MockPostServiceInterface, id int64) {
				s.EXPECT().Delete(id).Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"ok":true}`,
		},
		{
			name:      "Invalid id",
			postID:    "3f",
			postIDInt: 0,
			mockBehaviour: func(s *mock_post.MockPostServiceInterface, id int64) {
				s.EXPECT().Delete(id).Return(errors.New("test error")).AnyTimes()
			},
			expectedStatus: 400,
			expectedBody:   `{"message":"Invalid id"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockservice := mock_post.NewMockPostServiceInterface(c)
			test.mockBehaviour(mockservice, test.postIDInt)

			services := &service.BlogPostService{Post: mockservice}

			h := PostHandler{service: services.Post}

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.DELETE("/api/post/:id", h.HandlerDelete)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"DELETE",
				fmt.Sprintf("/api/post/%s", test.postID),
				bytes.NewBufferString(""),
			)
			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatus)
			assert.Equal(t, w.Body.String(), test.expectedBody)
		})
	}
}
