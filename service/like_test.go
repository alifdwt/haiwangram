package service

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/like"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/stretchr/testify/require"
)

func createRandomLike(t *testing.T) responses.LikeResponse {
	photo := createRandomPhoto(t)

	likeArg := &like.CreateLikeRequest{
		PhotoID: photo.ID,
	}

	err := likeArg.Validate()
	require.NoError(t, err)

	like, err := testService.Like.CreateLike(photo.UserID, *likeArg)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.NotZero(t, like.ID)
	require.Equal(t, photo.ID, like.PhotoID)
	require.Equal(t, photo.UserID, like.UserID)

	// Error not found
	_, err = testService.Like.CreateLike(0, *likeArg)
	require.Error(t, err)

	return *like
}

func TestCreateLike(t *testing.T) {
	createRandomLike(t)
}

func TestGetLikeById(t *testing.T) {
	like := createRandomLike(t)

	res, err := testService.Like.GetLikeById(like.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, like.ID, res.ID)
	require.Equal(t, like.PhotoID, res.PhotoID)
	require.Equal(t, like.UserID, res.UserID)
	require.NotZero(t, res.ID)

	// Error not found
	_, err = testService.Like.GetLikeById(0)
	require.Error(t, err)
}

func TestDeleteLike(t *testing.T) {
	like1 := createRandomLike(t)

	_, err := testService.Like.DeleteLike(like1.PhotoID, like1.UserID)
	require.NoError(t, err)

	res, err := testService.Like.GetLikeById(like1.ID)
	require.Empty(t, res)
	require.Error(t, err)

	// Error not found
	_, err = testService.Like.DeleteLike(0, 0)
	require.Error(t, err)
}
