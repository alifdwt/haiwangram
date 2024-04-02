package service

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/user"
	"github.com/alifdwt/haiwangram/util"
	"github.com/stretchr/testify/require"
)

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

	user2, err := testService.User.UpdateUserById(user1.ID, userArg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, fullName, user2.FullName)
	require.Equal(t, username, user2.Username)
	require.Equal(t, email, user2.Email)
	// require.Equal(t, birthDate, user2.BirthDate)
	require.Equal(t, profilePictureUrl, user2.ProfileImageURL)
	require.Equal(t, description, user2.Description)
	require.NotZero(t, user2.ID)

	require.NotEqual(t, user1.Email, user2.Email)
	require.NotEqual(t, user1.Username, user2.Username)
	require.NotEqual(t, user1.FullName, user2.FullName)
	// require.NotEqual(t, user1.BirthDate, user2.BirthDate)
	require.NotEqual(t, user1.ProfileImageURL, user2.ProfileImageURL)
	require.NotEqual(t, user1.Description, user2.Description)

	// Error not found
	_, err = testService.User.UpdateUserById(0, userArg)
	require.Error(t, err)
}

func TestDeleteUserById(t *testing.T) {
	user1 := createRandomUser(t)

	_, err := testService.User.DeleteUserById(user1.ID)
	require.NoError(t, err)

	// Error not found
	_, err = testService.User.DeleteUserById(0)
	require.Error(t, err)
}

func TestGetRandomUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	count := util.RandomInt(1, 10)
	res, err := testService.User.GetRandomUser(int(count))
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, int(count), len(*res))

	for _, user := range *res {
		require.NotEmpty(t, user)
	}
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testService.User.GetUserByEmail(user1.Email)
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
	_, err = testService.User.GetUserByEmail(util.RandomEmail(util.RandomOwner()))
	require.Error(t, err)
}

func TestGetUserByUsername(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testService.User.GetUserByUsername(user1.Username)
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
	_, err = testService.User.GetUserByUsername(util.RandomUsername(util.RandomOwner()))
	require.Error(t, err)
}
