package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/srjchsv/service/internal/handler"
	"github.com/srjchsv/service/internal/repository"
	"github.com/srjchsv/service/internal/services"
	mock_services "github.com/srjchsv/service/tests/internal/services/mock"
	"github.com/stretchr/testify/require"
)

type Cookie struct {
	Name  string
	Value string
}

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
		{
			name:      "Wrong input",
			inputBody: `{"username":"username", "name":"name"}`,
			inputUser: repository.User{
				Name:     "name",
				Username: "username",
				Password: "123",
			},
			mockBehavior:         func(r *mock_services.MockAuthorization, user repository.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"username":"username", "name":"name", "password":"123"}`,
			inputUser: repository.User{
				Name:     "name",
				Username: "username",
				Password: "123",
			},
			mockBehavior: func(r *mock_services.MockAuthorization, user repository.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
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
			handler := handler.NewHandler(services)

			//Init endpoint
			r := gin.New()
			handler.InitRouter(r)

			//Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			require.Equal(t, test.expectedResponseBody, w.Body.String())
			require.Equal(t, test.expectedStatusCode, w.Code)
		})

	}
}

func TestHandler_signIn(t *testing.T) {
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
			name:      "ok",
			inputBody: `{"username":"username","password":"password"}`,
			inputUser: repository.User{
				Username: "username",
				Password: "password",
			},
			mockBehavior: func(r *mock_services.MockAuthorization, user repository.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("token", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"token"}`,
		},
		{
			name:      "Wrong input",
			inputBody: `{"username":"username"}`,
			inputUser: repository.User{
				Username: "username",
				Password: "password",
			},
			mockBehavior:         func(r *mock_services.MockAuthorization, user repository.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "service error",
			inputBody: `{"username":"username","password":"password"}`,
			inputUser: repository.User{
				Username: "username",
				Password: "password",
			},
			mockBehavior: func(r *mock_services.MockAuthorization, user repository.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("0", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Init dependancies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_services.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)
			services := &services.Service{Authorization: repo}
			handler := handler.NewHandler(services)

			//Init endpoint
			r := gin.New()
			handler.InitRouter(r)

			//Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			require.Equal(t, test.expectedStatusCode, w.Code)
			require.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_refreshToken(t *testing.T) {
	type mockBehavior func(r *mock_services.MockAuthorization)

	tests := []struct {
		name                 string
		cookie               Cookie
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "ok",
			cookie: Cookie{Name: "access_token", Value: "token"},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("token").Return(1, nil)
				r.EXPECT().RefreshToken("token", 1).Return("newToken", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"newToken"}`,
		},
		{
			name:   "no cookie",
			cookie: Cookie{Name: "", Value: ""},
			mockBehavior: func(r *mock_services.MockAuthorization) {
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"no cookie access_token"}`,
		},
		{
			name:   "no token",
			cookie: Cookie{Name: "access_token", Value: ""},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("").Return(0, errors.New("no token"))

			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"no token"}`,
		},
		{
			name:   "not matching id",
			cookie: Cookie{Name: "access_token", Value: "token"},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("token").Return(1, nil)
				r.EXPECT().RefreshToken("token", 1).Return("", errors.New("cant refresh"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"cant refresh"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Init dependancies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_services.NewMockAuthorization(c)
			test.mockBehavior(repo)
			services := &services.Service{Authorization: repo}
			handler := handler.NewHandler(services)

			//Init endpoint
			r := gin.New()
			gin.SetMode(gin.ReleaseMode)
			handler.InitRouter(r)

			//Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/refresh-token", nil)
			req.AddCookie(&http.Cookie{Name: test.cookie.Name, Value: test.cookie.Value})
			r.ServeHTTP(w, req)

			require.Equal(t, test.expectedStatusCode, w.Code)
			require.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_logout(t *testing.T) {
	type mockBehavior func(r *mock_services.MockAuthorization)
	tests := []struct {
		name                 string
		cookie               Cookie
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "ok",
			cookie: Cookie{Name: "access_token", Value: "token"},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("token").Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"logout":"success"}`,
		},
		{
			name:   "no cookie",
			cookie: Cookie{Name: "", Value: ""},
			mockBehavior: func(r *mock_services.MockAuthorization) {
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"no cookie access_token"}`,
		},
		{
			name:   "no token",
			cookie: Cookie{Name: "access_token", Value: ""},
			mockBehavior: func(r *mock_services.MockAuthorization) {
				r.EXPECT().ParseToken("").Return(0, errors.New("no token"))

			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"no token"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Init dependancies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_services.NewMockAuthorization(c)
			test.mockBehavior(repo)
			services := &services.Service{Authorization: repo}
			handler := handler.NewHandler(services)

			//Init endpoint
			r := gin.New()
			gin.SetMode(gin.ReleaseMode)
			handler.InitRouter(r)

			//Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/logout", nil)
			req.AddCookie(&http.Cookie{Name: test.cookie.Name, Value: test.cookie.Value})
			r.ServeHTTP(w, req)

			require.Equal(t, test.expectedStatusCode, w.Code)
			require.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}
