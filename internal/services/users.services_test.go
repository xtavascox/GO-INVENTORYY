package services

import (
	"context"
	"github.com/go-inventory/encryption"
	"github.com/go-inventory/internal/entity"
	"github.com/go-inventory/internal/repository"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)

	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(&entity.User{Email: "test@exists.com", Password: encryptedPassword}, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {

	testCases := []struct {
		Name          string
		Email         string
		Password      string
		UserName      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "test@test.com",
			Password:      "validPassword",
			UserName:      "test",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "test@exists.com",
			Password:      "validPassword",
			UserName:      "test",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			s := New(repo)

			err := s.RegisterUser(ctx, tc.Email, tc.Password, tc.UserName)
			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}

		})
	}

}

func TestLoginUser(t *testing.T) {

	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_InvalidCredentials",
			Email:         "test@exists.com",
			Password:      "invalidPassword",
			ExpectedError: ErrInvalidCredentials,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}

		})
	}
}
