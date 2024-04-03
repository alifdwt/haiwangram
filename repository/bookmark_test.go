package repository

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/bookmark"
	"github.com/alifdwt/haiwangram/models"
	"github.com/stretchr/testify/require"
)

func createRandomBookmark(t *testing.T) *models.Bookmark {
	photo := createRandomPhoto(t)

	bookmarkArg := &bookmark.CreateBookmarkRequest{
		PhotoID: photo.ID,
	}

	err := bookmarkArg.Validate()
	require.NoError(t, err)

	bookmark, err := testRepository.Bookmark.CreateBookmark(photo.UserID, *bookmarkArg)
	require.NoError(t, err)
	require.NotEmpty(t, bookmark)

	require.NotZero(t, bookmark.ID)
	require.Equal(t, photo.ID, bookmark.PhotoID)
	require.Equal(t, photo.UserID, bookmark.UserID)

	// Error not found
	_, err = testRepository.Bookmark.CreateBookmark(0, *bookmarkArg)
	require.Error(t, err)

	return bookmark
}

func TestCreateBookmark(t *testing.T) {
	createRandomBookmark(t)
}

func TestGetBookmarkById(t *testing.T) {
	bookmark := createRandomBookmark(t)

	res, err := testRepository.Bookmark.GetBookmarkById(bookmark.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, bookmark.ID, res.ID)
	require.Equal(t, bookmark.PhotoID, res.PhotoID)
	require.Equal(t, bookmark.UserID, res.UserID)
	require.NotZero(t, res.ID)

	// Error not found
	_, err = testRepository.Bookmark.GetBookmarkById(0)
	require.Error(t, err)
}

func TestGetBookmarkByUserId(t *testing.T) {
	bookmark := createRandomBookmark(t)

	bookmarks, err := testRepository.Bookmark.GetBookmarkByUserId(bookmark.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, bookmarks)

	for _, bookmark := range *bookmarks {
		require.NotEmpty(t, bookmark)
		require.Equal(t, bookmark.UserID, bookmark.UserID)
	}
}

func TestDeleteBookmark(t *testing.T) {
	bookmark1 := createRandomBookmark(t)

	_, err := testRepository.Bookmark.DeleteBookmark(bookmark1.PhotoID, bookmark1.UserID)
	require.NoError(t, err)

	bookmark2, err := testRepository.Bookmark.GetBookmarkById(bookmark1.ID)
	require.Empty(t, bookmark2)
	require.Error(t, err)

	// Error not found
	_, err = testRepository.Bookmark.DeleteBookmark(0, 0)
	require.Error(t, err)
}
