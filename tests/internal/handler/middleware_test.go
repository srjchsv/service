package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/srjchsv/service/internal/handler"
	"github.com/srjchsv/service/internal/services"
	mock_services "github.com/srjchsv/service/tests/internal/services/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(r *mock_services.MockAuthorization)
	type cookie struct {
		Name  string
		Value string
	}
	tests := []struct {
		name               string
		cookie             cookie
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name: "ok",
			cookie: cookie{
				Name:  "access_token",
				Value: "token",
			},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("token").Return(1, nil)
			},
			expectedStatusCode: 200,
		},
		{
			name: "no cookie",
			cookie: cookie{
				Name:  "",
				Value: "",
			},
			mockBehavior:       func(r *mock_services.MockAuthorization) {},
			expectedStatusCode: 401,
		},
		{
			name: "error parsing token",
			cookie: cookie{
				Name:  "access_token",
				Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA1NTAyMDIsImlhdCI6MTY2MDUwNzAwMiwidXNlcl9pZCI6MX0",
			},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA1NTAyMDIsImlhdCI6MTY2MDUwNzAwMiwidXNlcl9pZCI6MX0").Return(0, errors.New("ERROR"))
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
		handler := handler.NewHandler(services)

		r := gin.New()
		handler.InitRouter(r)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api", nil)

		cookie := &http.Cookie{
			Name:  test.cookie.Name,
			Value: test.cookie.Value,
		}

		req.AddCookie(cookie)
		r.ServeHTTP(w, req)
		w.Header()

		require.Equal(t, test.expectedStatusCode, w.Code)
	}
}
