package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/srjchsv/service/internal/repository"
	mock_repository "github.com/srjchsv/service/internal/repository/mock"
	"github.com/stretchr/testify/require"
)

func TestService_GenerateToken(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockAuthorization, user repository.User)

	tests := []struct {
		name             string
		inputUser        repository.User
		mockBehavior     mockBehavior
		expectedResponse error
	}{
		{
			name: "ok",
			inputUser: repository.User{
				Username: "username",
				Password: "password",
			},
			mockBehavior: func(r *mock_repository.MockAuthorization, user repository.User) {
				r.EXPECT().GetUser("username", "6673666466333434343464696a69736a64666a6466695baa61e4c9b93f3f0682250b6cf8331b7ee68fd8").Return(repository.User{
					ID: 1,
				}, nil)
			},
			expectedResponse: nil,
		},
		{
			name: "error",
			inputUser: repository.User{
				Password: "password",
			},
			mockBehavior: func(r *mock_repository.MockAuthorization, user repository.User) {
				r.EXPECT().GetUser("", "6673666466333434343464696a69736a64666a6466695baa61e4c9b93f3f0682250b6cf8331b7ee68fd8").Return(repository.User{
					ID: 1,
				}, errors.New("no username"))
			},
			expectedResponse: errors.New("no username"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			s := NewAuthService(repo)
			_, err := s.GenerateToken(test.inputUser.Username, test.inputUser.Password)

			require.Equal(t, test.expectedResponse, err)
		})
	}
}

func TestService_ParseToken(t *testing.T) {

	tests := []struct {
		name          string
		inputToken    string
		expected      int
		expectedError error
	}{
		{
			name:          "ok",
			inputToken:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA1NTAyMDIsImlhdCI6MTY2MDUwNzAwMiwidXNlcl9pZCI6MX0.PJtZzvUjY6IHOtAUOiesLgJ2Sft0JkOBM-_2pUf5duw",
			expected:      1,
			expectedError: nil,
		},
		{
			name:          "missing token",
			inputToken:    "",
			expected:      0,
			expectedError: errors.New("token contains an invalid number of segments"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockAuthorization(c)
			s := NewAuthService(repo)
			id, err := s.ParseToken(test.inputToken)

			require.Equal(t, test.expectedError, err)
			require.Equal(t, test.expected, id)
		})
	}
}
