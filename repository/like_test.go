package repository

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/like"
	"github.com/alifdwt/haiwangram/models"
	"github.com/stretchr/testify/require"
)

func createRandomLike(t *testing.T) *models.Like {
	photo := createRandomPhoto(t)

	likeArg := &like.CreateLikeRequest{
		PhotoID: photo.ID,
	}

	err := likeArg.Validate()
	require.NoError(t, err)

	like, err := testRepository.Like.CreateLike(photo.UserID, *likeArg)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.NotZero(t, like.ID)
	require.Equal(t, photo.ID, like.PhotoID)
	require.Equal(t, photo.UserID, like.UserID)

	// Error not found
	_, err = testRepository.Like.CreateLike(0, *likeArg)
	require.Error(t, err)

	return like
}

func TestCreateLike(t *testing.T) {
	createRandomLike(t)
}

func TestGetLikeById(t *testing.T) {
	like := createRandomLike(t)

	res, err := testRepository.Like.GetLikeById(like.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, like.ID, res.ID)
	require.Equal(t, like.PhotoID, res.PhotoID)
	require.Equal(t, like.UserID, res.UserID)
	require.NotZero(t, res.ID)

	// Error not found
	_, err = testRepository.Like.GetLikeById(0)
	require.Error(t, err)
}

func TestDeleteLike(t *testing.T) {
	like1 := createRandomLike(t)

	_, err := testRepository.Like.DeleteLike(like1.PhotoID, like1.UserID)
	require.NoError(t, err)

	res, err := testRepository.Like.GetLikeById(like1.ID)
	require.Empty(t, res)
	require.Error(t, err)

	// Error not found
	_, err = testRepository.Like.DeleteLike(0, 0)
	require.Error(t, err)
}
