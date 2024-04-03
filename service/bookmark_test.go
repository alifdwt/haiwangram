package service

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/bookmark"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/stretchr/testify/require"
)

func createRandomBookmark(t *testing.T) *responses.BookmarkResponse {
	photo := createRandomPhoto(t)

	bookmarkArg := &bookmark.CreateBookmarkRequest{
		PhotoID: photo.ID,
	}

	err := bookmarkArg.Validate()
	require.NoError(t, err)

	bookmark, err := testService.Bookmark.CreateBookmark(photo.UserID, *bookmarkArg)
	require.NoError(t, err)
	require.NotEmpty(t, bookmark)

	require.NotZero(t, bookmark.ID)
	require.Equal(t, photo.ID, bookmark.PhotoID)
	require.Equal(t, photo.UserID, bookmark.UserID)

	// Error not found
	_, err = testService.Bookmark.CreateBookmark(0, *bookmarkArg)
	require.Error(t, err)

	return bookmark
}

func TestCreateBookmark(t *testing.T) {
	createRandomBookmark(t)
}

func TestGetBookmarkById(t *testing.T) {
	bookmark := createRandomBookmark(t)

	res, err := testService.Bookmark.GetBookmarkById(bookmark.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, bookmark.ID, res.ID)
	require.Equal(t, bookmark.PhotoID, res.PhotoID)
	require.Equal(t, bookmark.UserID, res.UserID)
	require.NotZero(t, res.ID)

	// Error not found
	_, err = testService.Bookmark.GetBookmarkById(0)
	require.Error(t, err)
}

func TestGetBookmarkByUserId(t *testing.T) {
	bookmark := createRandomBookmark(t)

	bookmarks, err := testService.Bookmark.GetBookmarkByUserId(bookmark.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, bookmarks)

	for _, bookmark := range *bookmarks {
		require.NotEmpty(t, bookmark)
		require.Equal(t, bookmark.UserID, bookmark.UserID)
	}
}

func TestDeleteBookmark(t *testing.T) {
	bookmark1 := createRandomBookmark(t)

	_, err := testService.Bookmark.DeleteBookmark(bookmark1.PhotoID, bookmark1.UserID)
	require.NoError(t, err)

	bookmark2, err := testService.Bookmark.GetBookmarkById(bookmark1.ID)
	require.Empty(t, bookmark2)
	require.Error(t, err)

	// Error not found
	_, err = testService.Bookmark.DeleteBookmark(0, 0)
	require.Error(t, err)
}
