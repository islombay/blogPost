package post_handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/islombay/blogPost/internal/service"
	"github.com/islombay/blogPost/internal/service/post"
	mock_post "github.com/islombay/blogPost/internal/service/post/mocks"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPostHandler_HandlerGetAll(t *testing.T) {
	type mockBehaviour func(s *mock_post.MockPostServiceInterface)

	testTable := []struct {
		name           string
		mockBehaviour  mockBehaviour
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_post.MockPostServiceInterface) {
				r := []post.PostModel{
					post.PostModel{
						ID:        1,
						Title:     "Title",
						Content:   "Content",
						CreatedAt: time.Date(2023, 11, 28, 15, 5, 24, 0, time.UTC),
						Username:  "islombay",
					},
					post.PostModel{
						ID:        2,
						Title:     "Title2",
						Content:   "Content2",
						CreatedAt: time.Date(2023, 11, 28, 15, 5, 24, 0, time.UTC),
					},
				}
				s.EXPECT().GetAll().Return(r, nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"posts":[{"id":1,"title":"Title","content":"Content","created_at":"2023-11-28T15:05:24Z","username":"islombay"},{"id":2,"title":"Title2","content":"Content2","created_at":"2023-11-28T15:05:24Z"}]}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockservice := mock_post.NewMockPostServiceInterface(c)
			test.mockBehaviour(mockservice)

			services := &service.BlogPostService{Post: mockservice}
			post_handler := PostHandler{service: services.Post}

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/api/post/all", post_handler.HandlerGetAll)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"GET",
				"/api/post/all",
				bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			// Check results
			assert.Equal(t, w.Code, test.expectedStatus)
			assert.Equal(t, w.Body.String(), test.expectedBody)
		})
	}
}
