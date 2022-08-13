package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/srjchsv/service/internal/repository"
	"github.com/srjchsv/service/internal/services"
	mock_services "github.com/srjchsv/service/internal/services/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *mock_services.MockAuthorization, user repository.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            repository.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username":"username", "name":"name", "password":"123"}`,
			inputUser: repository.User{
				Name:     "name",
				Username: "username",
				Password: "123",
			},
			mockBehavior: func(r *mock_services.MockAuthorization, user repository.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_services.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &services.Service{Authorization: repo}
			handler := Handler{services}

			//Init endpoint
			r := gin.New()

			r.POST("/sign-up", handler.signUp)

			//Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			require.Equal(t, test.expectedStatusCode, w.Code)
			require.Equal(t, test.expectedResponseBody, w.Body.String())
		})

	}
}
