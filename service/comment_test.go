package service

import (
	"testing"

	"github.com/alifdwt/haiwangram/domain/requests/comment"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/util"
	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T) responses.CommentResponse {
	photo := createRandomPhoto(t)
	message := util.RandomWords(7)
	photoId := photo.ID

	commentArg := &comment.CreateCommentRequest{
		Message: message,
		PhotoID: photoId,
	}

	err := commentArg.Validate()
	require.NoError(t, err)

	comment, err := testService.Comment.CreateComment(photo.UserID, *commentArg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.NotZero(t, comment.ID)
	require.Equal(t, message, comment.Message)
	require.Equal(t, photoId, comment.PhotoID)
	require.Equal(t, photo.UserID, comment.UserID)

	// Error not found
	_, err = testService.Comment.CreateComment(0, *commentArg)
	require.Error(t, err)

	return *comment
}

func TestCreateComment(t *testing.T) {
	createRandomComment(t)
}

func TestGetCommentAll(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomComment(t)
	}

	res, err := testService.Comment.GetCommentAll()
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, comment := range *res {
		require.NotEmpty(t, comment)
	}
}

func TestGetCommentById(t *testing.T) {
	comment := createRandomComment(t)

	res, err := testService.Comment.GetCommentById(comment.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, comment.ID, res.ID)
	require.Equal(t, comment.Message, res.Message)
	require.Equal(t, comment.PhotoID, res.PhotoID)
	require.Equal(t, comment.UserID, res.UserID)

	require.NotZero(t, res.ID)

	// Error not found
	_, err = testService.Comment.GetCommentById(0)
	require.Error(t, err)
}

func TestGetCommentByPhotoId(t *testing.T) {
	newComment := createRandomComment(t)

	comments, err := testService.Comment.GetCommentByPhotoId(newComment.PhotoID)
	require.NoError(t, err)
	require.NotEmpty(t, comments)

	for _, comment := range *comments {
		require.NotEmpty(t, comment)
		require.Equal(t, newComment.PhotoID, comment.PhotoID)
	}
}

func TestUpdateComment(t *testing.T) {
	comment1 := createRandomComment(t)

	message := util.RandomWords(5)

	commentArg := &comment.UpdateCommentRequest{
		Message: message,
	}

	err := commentArg.Validate()
	require.NoError(t, err)

	comment2, err := testService.Comment.UpdateComment(comment1.ID, *commentArg)
	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, message, comment2.Message)
	require.Equal(t, comment1.PhotoID, comment2.PhotoID)
	require.Equal(t, comment1.UserID, comment2.UserID)

	require.NotEqual(t, comment1.Message, comment2.Message)

	// Error not found
	_, err = testService.Comment.UpdateComment(0, *commentArg)
	require.Error(t, err)
}

func TestDeleteComment(t *testing.T) {
	comment1 := createRandomComment(t)

	_, err := testService.Comment.DeleteComment(comment1.ID)
	require.NoError(t, err)

	_, err = testService.Comment.GetCommentById(comment1.ID)
	require.Error(t, err)

	// Error not found
	_, err = testService.Comment.DeleteComment(0)
	require.Error(t, err)
}
