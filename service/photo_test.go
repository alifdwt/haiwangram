package service

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/photo"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/util"
	"github.com/stretchr/testify/require"
)

func createRandomPhoto(t *testing.T) responses.PhotoResponse {
	user := createRandomUser(t)
	title := util.RandomWords(7)
	caption := util.RandomWords(15)
	photoUrl := util.RandomProfilePictureUrl()
	userId := user.ID

	photoArg := &photo.CreatePhotoRequest{
		Title:    title,
		Caption:  caption,
		PhotoURL: photoUrl,
	}

	err := photoArg.Validate()
	require.NoError(t, err)

	photo, err := testService.Photo.CreatePhoto(userId, *photoArg)
	require.NoError(t, err)
	require.NotEmpty(t, photo)

	require.NotZero(t, photo.ID)
	require.Equal(t, title, photo.Title)
	require.Equal(t, caption, photo.Caption)
	require.Equal(t, photoUrl, photo.PhotoURL)
	require.Equal(t, userId, photo.UserID)

	// Error not found
	_, err = testService.Photo.CreatePhoto(0, *photoArg)
	require.Error(t, err)

	return *photo
}

func TestCreatePhoto(t *testing.T) {
	createRandomPhoto(t)
}

func TestGetPhotoAll(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPhoto(t)
	}

	photos, err := testService.Photo.GetPhotoAll()
	require.NoError(t, err)
	require.NotEmpty(t, photos)

	for _, photo := range *photos {
		require.NotEmpty(t, photo)
	}
}

func TestGetPhotoById(t *testing.T) {
	photo := createRandomPhoto(t)

	res, err := testService.Photo.GetPhotoById(photo.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, photo.ID, res.ID)
	require.Equal(t, photo.Title, res.Title)
	require.Equal(t, photo.Caption, res.Caption)
	require.Equal(t, photo.PhotoURL, res.PhotoURL)
	require.Equal(t, photo.UserID, res.UserID)

	require.NotZero(t, res.ID)

	// Error not found
	_, err = testService.Photo.GetPhotoById(0)
	require.Error(t, err)
}

func TestUpdatePhoto(t *testing.T) {
	photo1 := createRandomPhoto(t)

	title := util.RandomWords(5)
	caption := util.RandomWords(10)
	photoUrl := util.RandomProfilePictureUrl()

	photoArg := &photo.UpdatePhotoRequest{
		Title:    title,
		Caption:  caption,
		PhotoURL: photoUrl,
	}

	err := photoArg.Validate()
	require.NoError(t, err)

	photo2, err := testService.Photo.UpdatePhoto(photo1.UserID, photo1.ID, *photoArg)
	require.NoError(t, err)
	require.NotEmpty(t, photo2)

	require.Equal(t, title, photo2.Title)
	require.Equal(t, caption, photo2.Caption)
	require.Equal(t, photoUrl, photo2.PhotoURL)
	require.Equal(t, photo1.UserID, photo2.UserID)
	require.Equal(t, photo1.ID, photo2.ID)

	require.NotEqual(t, photo1.PhotoURL, photo2.PhotoURL)
	require.NotEqual(t, photo1.Title, photo2.Title)
	require.NotEqual(t, photo1.Caption, photo2.Caption)

	// Error not found
	_, err = testService.Photo.UpdatePhoto(0, 0, *photoArg)
	require.Error(t, err)
}

func TestDeletePhoto(t *testing.T) {
	photo1 := createRandomPhoto(t)

	_, err := testService.Photo.DeletePhoto(photo1.ID)
	require.NoError(t, err)

	_, err = testService.Photo.GetPhotoById(photo1.ID)
	require.Error(t, err)

	// Error not found
	_, err = testService.Photo.DeletePhoto(0)
	require.Error(t, err)
}
