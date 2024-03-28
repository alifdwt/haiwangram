package service

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/auth"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) responses.UserResponse {
	fullName := util.RandomOwner()
	username := util.RandomUsername(fullName)
	email := util.RandomEmail(username)
	birthDate := util.RandomBirthDate(int(util.RandomInt(20, 30)))
	profilePictureUrl := util.RandomProfilePictureUrl()
	password := "123456"
	description := util.RandomWords(7)

	userArg := &auth.RegisterRequest{
		FullName:        fullName,
		Username:        username,
		Email:           email,
		BirthDate:       birthDate,
		ProfileImageURL: profilePictureUrl,
		Password:        password,
		Description:     description,
	}

	user, err := testService.Auth.Register(userArg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, fullName, user.FullName)
	require.Equal(t, username, user.Username)
	require.Equal(t, email, user.Email)
	// require.Equal(t, birthDate, user.BirthDate)
	require.Equal(t, profilePictureUrl, user.ProfileImageURL)
	require.Equal(t, description, user.Description)
	require.NotZero(t, user.ID)

	return *user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestLogin(t *testing.T) {
	user := createRandomUser(t)

	loginArg := &auth.LoginRequest{
		Email:    user.Email,
		Password: "123456",
	}

	token, err := testService.Auth.Login(loginArg)
	require.NoError(t, err)
	require.NotEmpty(t, token)
}
