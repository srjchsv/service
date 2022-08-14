package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/srjchsv/service/internal/services"
	mock_services "github.com/srjchsv/service/internal/services/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(r *mock_services.MockAuthorization)
	tests := []struct {
		name               string
		header             map[string][]string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name: "ok",
			header: map[string][]string{
				"Accept":        {"application/json"},
				"Authorization": {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA1NTAyMDIsImlhdCI6MTY2MDUwNzAwMiwidXNlcl9pZCI6MX0.PJtZzvUjY6IHOtAUOiesLgJ2Sft0JkOBM-_2pUf5duw"},
			},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA1NTAyMDIsImlhdCI6MTY2MDUwNzAwMiwidXNlcl9pZCI6MX0.PJtZzvUjY6IHOtAUOiesLgJ2Sft0JkOBM-_2pUf5duw").Return(1, nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:               "empty auth header",
			header:             map[string][]string{},
			mockBehavior:       func(r *mock_services.MockAuthorization) {},
			expectedStatusCode: 401,
		},
		{
			name: "invalid auth header",
			header: map[string][]string{
				"Accept":        {"application/json"},
				"Authorization": {"test test "},
			},
			mockBehavior:       func(r *mock_services.MockAuthorization) {},
			expectedStatusCode: 401,
		},
		{
			name: "token is empty",
			header: map[string][]string{
				"Accept":        {"application/json"},
				"Authorization": {"Bearer "},
			},
			mockBehavior: func(r *mock_services.MockAuthorization) {
			},
			expectedStatusCode: 401,
		},
		{
			name: "invalid token",
			header: map[string][]string{
				"Accept":        {"application/json"},
				"Authorization": {"Bearer gfdgdsfdsfdsf"},
			},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("gfdgdsfdsfdsf").Return(0, errors.New("token contains an invalid number of segments"))
			},
			expectedStatusCode: 401,
		},
	}
	for _, test := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := mock_services.NewMockAuthorization(c)
		test.mockBehavior(repo)
		services := &services.Service{Authorization: repo}
		handler := &Handler{services: services}

		r := gin.New()
		apiV1 := r.Group("/api", handler.userIdentity)
		apiV1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hi you are in the secured route...",
			})
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/", nil)
		req.Header = test.header
		r.ServeHTTP(w, req)
		w.Header()

		require.Equal(t, test.expectedStatusCode, w.Code)
	}
}
