package repository

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/auth"
	"github.com/alifdwt/haiwangram/domain/requests/user"
	"github.com/alifdwt/haiwangram/models"
	"github.com/alifdwt/haiwangram/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *models.User {
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

	err := userArg.Validate()
	require.NoError(t, err)

	user, err := testRepository.User.CreateUser(userArg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.FullName, fullName)
	require.Equal(t, user.Username, username)
	require.Equal(t, user.Email, email)
	require.Equal(t, user.BirthDate, birthDate)
	require.Equal(t, user.ProfileImageURL, profilePictureUrl)
	require.Equal(t, user.Description, description)
	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testRepository.User.GetUserByEmail(user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.BirthDate.Format("2006-01-02"), user2.BirthDate.Format("2006-01-02"))
	require.Equal(t, user1.ProfileImageURL, user2.ProfileImageURL)
	require.Equal(t, user1.Description, user2.Description)

	require.Equal(t, user1.ID, user2.ID)

	// Error not found
	_, err = testRepository.User.GetUserByEmail(util.RandomEmail(util.RandomOwner()))
	require.Error(t, err)
}

func TestGetUserById(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testRepository.User.GetUserById(user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.BirthDate.Format("2006-01-02"), user2.BirthDate.Format("2006-01-02"))
	require.Equal(t, user1.ProfileImageURL, user2.ProfileImageURL)
	require.Equal(t, user1.Description, user2.Description)

	require.Equal(t, user1.ID, user2.ID)

	// Error not found
	_, err = testRepository.User.GetUserById(0)
	require.Error(t, err)
}

func TestUpdateUserById(t *testing.T) {
	user1 := createRandomUser(t)

	fullName := util.RandomOwner()
	username := util.RandomUsername(fullName)
	email := util.RandomEmail(username)
	birthDate := util.RandomBirthDate(int(util.RandomInt(20, 30)))
	profilePictureUrl := util.RandomProfilePictureUrl()
	password := "12345678"
	description := util.RandomWords(5)

	userArg := &user.UpdateUserRequest{
		Email:           email,
		Username:        username,
		FullName:        fullName,
		BirthDate:       birthDate,
		ProfileImageURL: profilePictureUrl,
		Password:        password,
		Description:     description,
	}

	err := userArg.Validate()
	require.NoError(t, err)

	user2, err := testRepository.User.UpdateUserById(user1.ID, userArg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, fullName, user2.FullName)
	require.Equal(t, username, user2.Username)
	require.Equal(t, email, user2.Email)
	require.Equal(t, birthDate, user2.BirthDate)
	require.Equal(t, profilePictureUrl, user2.ProfileImageURL)
	require.Equal(t, description, user2.Description)
	require.NotZero(t, user2.ID)

	require.NotEqual(t, user1.Email, user2.Email)
	require.NotEqual(t, user1.Username, user2.Username)
	require.NotEqual(t, user1.FullName, user2.FullName)
	require.NotEqual(t, user1.BirthDate, user2.BirthDate)
	require.NotEqual(t, user1.ProfileImageURL, user2.ProfileImageURL)
	require.NotEqual(t, user1.Description, user2.Description)

	require.Equal(t, password, user2.Password)
	require.NotEqual(t, user1.Password, user2.Password)

	// Error not found
	_, err = testRepository.User.UpdateUserById(0, userArg)
	require.Error(t, err)
}

func TestDeleteUserById(t *testing.T) {
	user1 := createRandomUser(t)

	_, err := testRepository.User.DeleteUserById(user1.ID)
	require.NoError(t, err)

	// Error not found
	_, err = testRepository.User.DeleteUserById(0)
	require.Error(t, err)
}
