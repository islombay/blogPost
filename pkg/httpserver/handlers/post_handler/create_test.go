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

func TestPostHandler_HandlerCreate(t *testing.T) {
	type mockBehaviour func(s *mock_post.MockPostServiceInterface, model post.CreateModel)

	testTable := []struct {
		name            string
		inputBody       string
		inputPostModel  post.CreateModel
		mockBehaviour   mockBehaviour
		expectedStatus  int
		expectedReqBody string
	}{
		{
			name: "OK without username",
			inputBody: `{"title": "Test case for db with success",
						"content": "Test content OK",
						"created_at": "2023-11-28T15:05:24Z"}`,
			inputPostModel: post.CreateModel{
				Title:     "Test case for db with success",
				Content:   "Test content OK",
				CreatedAt: time.Date(2023, 11, 28, 15, 5, 24, 0, time.UTC),
			},
			mockBehaviour: func(s *mock_post.MockPostServiceInterface, model post.CreateModel) {
				response := post.PostModel{
					ID:        1,
					Title:     model.Title,
					Content:   model.Content,
					CreatedAt: model.CreatedAt,
					Username:  model.Username,
				}
				s.EXPECT().Create(model).Return(response, nil)
			},
			expectedStatus:  200,
			expectedReqBody: `{"id":1,"title":"Test case for db with success","content":"Test content OK","created_at":"2023-11-28T15:05:24Z"}`,
		},
		{
			name: "OK",
			inputBody: `{"title": "Test case for db with success and username",
						"content": "Test content OK",
						"created_at": "2023-11-28T15:05:24Z",
						"username":"tester"}`,
			inputPostModel: post.CreateModel{
				Title:     "Test case for db with success and username",
				Content:   "Test content OK",
				CreatedAt: time.Date(2023, 11, 28, 15, 5, 24, 0, time.UTC),
				Username:  "tester",
			},
			mockBehaviour: func(s *mock_post.MockPostServiceInterface, model post.CreateModel) {
				response := post.PostModel{
					ID:        1,
					Title:     model.Title,
					Content:   model.Content,
					CreatedAt: model.CreatedAt,
					Username:  model.Username,
				}
				s.EXPECT().Create(model).Return(response, nil)
			},
			expectedStatus:  200,
			expectedReqBody: `{"id":1,"title":"Test case for db with success and username","content":"Test content OK","created_at":"2023-11-28T15:05:24Z","username":"tester"}`,
		},
		{
			name:           "bad request in json",
			inputBody:      `{"title":"","content": "something", "created_at":"2023-11-28T15:05:24Z"}`,
			inputPostModel: post.CreateModel{},
			mockBehaviour: func(s *mock_post.MockPostServiceInterface, model post.CreateModel) {
				response := post.PostModel{
					ID:        1,
					Title:     model.Title,
					Content:   model.Content,
					CreatedAt: model.CreatedAt,
					Username:  model.Username,
				}
				s.EXPECT().Create(model).Return(response, nil).AnyTimes()
			},
			expectedStatus:  400,
			expectedReqBody: `{"message":"Post Body is not correct"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockservice := mock_post.NewMockPostServiceInterface(c)
			test.mockBehaviour(mockservice, test.inputPostModel)

			services := &service.BlogPostService{Post: mockservice}
			post_handler := PostHandler{service: services.Post}

			// Router init
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/api/post/new", post_handler.HandlerCreate)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"POST",
				"/api/post/new",
				bytes.NewBufferString(test.inputBody),
			)

			// Perform request
			r.ServeHTTP(w, req)

			// Check results
			assert.Equal(t, w.Code, test.expectedStatus)
			assert.Equal(t, w.Body.String(), test.expectedReqBody)
		})
	}
}
