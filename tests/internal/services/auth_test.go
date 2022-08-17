package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/srjchsv/service/internal/repository"
	"github.com/srjchsv/service/internal/services"
	mock_repository "github.com/srjchsv/service/tests/internal/repository/mock"
	"github.com/stretchr/testify/require"
)

const (
	userIDForTests = 1
)

func generateTokenForTest(t *testing.T) string {
	//init dependencies
	c := gomock.NewController(t)
	defer c.Finish()

	repo := mock_repository.NewMockAuthorization(c)

	repo.EXPECT().GetUser("username", "6673666466333434343464696a69736a64666a6466695baa61e4c9b93f3f0682250b6cf8331b7ee68fd8").Return(repository.User{
		ID: userIDForTests,
	}, nil)
	s := services.NewAuthService(repo)
	token, _ := s.GenerateToken("username", "password")

	return token
}

func TestService_GenerateToken(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockAuthorization, user repository.User)

	tests := []struct {
		name          string
		inputUser     repository.User
		mockBehavior  mockBehavior
		expectedError error
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
			expectedError: nil,
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
			expectedError: errors.New("no username"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			s := services.NewAuthService(repo)
			_, err := s.GenerateToken(test.inputUser.Username, test.inputUser.Password)

			require.Equal(t, test.expectedError, err)
		})
	}
}

func TestService_ParseToken(t *testing.T) {
	token := generateTokenForTest(t)

	tests := []struct {
		name          string
		inputToken    string
		expected      int
		expectedError error
	}{
		{
			name:          "ok",
			inputToken:    token,
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
			s := services.NewAuthService(repo)
			id, err := s.ParseToken(test.inputToken)

			require.Equal(t, test.expectedError, err)
			require.Equal(t, test.expected, id)
		})
	}
}

func TestService_RefreshToken(t *testing.T) {
	token := generateTokenForTest(t)
	tests := []struct {
		name             string
		currentToken     string
		currentUserID    int
		expectedError    bool
		expectedNewToken bool
	}{
		{
			name:             "ok",
			currentToken:     token,
			currentUserID:    1,
			expectedError:    false,
			expectedNewToken: true,
		},
		{
			name:             "missing token",
			currentToken:     "",
			currentUserID:    1,
			expectedError:    true,
			expectedNewToken: false,
		},
		{
			name:             "wrong user",
			currentToken:     token,
			currentUserID:    2,
			expectedError:    true,
			expectedNewToken: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockAuthorization(c)
			s := services.NewAuthService(repo)
			newToken, err := s.RefreshToken(test.currentToken, test.currentUserID)
			require.Equal(t, test.expectedNewToken, len(newToken) > 0)
			require.Equal(t, test.expectedError, err != nil)
		})
	}
}
