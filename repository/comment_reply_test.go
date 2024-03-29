package repository

import (
	"testing"

	commentreply "github.com/alifdwt/haiwangram/domain/requests/comment_reply"
	"github.com/alifdwt/haiwangram/models"
	"github.com/alifdwt/haiwangram/util"
	"github.com/stretchr/testify/require"
)

func createRandomCommentReply(t *testing.T) *models.CommentReply {
	comment := createRandomComment(t)
	message := util.RandomWords(10)

	commentReplyArg := &commentreply.CreateCommentReplyRequest{
		Message:   message,
		CommentID: comment.ID,
	}

	err := commentReplyArg.Validate()
	require.NoError(t, err)

	commentReply, err := testRepository.CommentReply.CreateCommentReply(comment.UserID, *commentReplyArg)
	require.NoError(t, err)
	require.NotEmpty(t, commentReply)

	require.NotZero(t, commentReply.ID)
	require.Equal(t, message, commentReply.Message)
	require.Equal(t, comment.ID, commentReply.CommentID)
	require.Equal(t, comment.UserID, commentReply.UserID)

	// Error not found
	_, err = testRepository.CommentReply.CreateCommentReply(0, *commentReplyArg)
	require.Error(t, err)

	return commentReply
}

func TestCreateCommentReply(t *testing.T) {
	createRandomCommentReply(t)
}

func TestGetComentReplyAll(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCommentReply(t)
	}

	commentReplies, err := testRepository.CommentReply.GetCommentReplyAll()
	require.NoError(t, err)
	require.NotEmpty(t, commentReplies)

	for _, commentReply := range *commentReplies {
		require.NotEmpty(t, commentReply)
	}
}

func TestGetCommentReplyById(t *testing.T) {
	commentReply := createRandomCommentReply(t)

	res, err := testRepository.CommentReply.GetCommentReplyById(commentReply.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, commentReply.ID, res.ID)
	require.Equal(t, commentReply.Message, res.Message)
	require.Equal(t, commentReply.CommentID, res.CommentID)
	require.Equal(t, commentReply.UserID, res.UserID)

	require.NotZero(t, res.ID)

	// Error not found
	_, err = testRepository.CommentReply.GetCommentReplyById(0)
	require.Error(t, err)
}

func TestUpdateCommentReply(t *testing.T) {
	commentReply1 := createRandomCommentReply(t)

	message := util.RandomWords(5)

	commentReplyArg := &commentreply.UpdateCommentReplyRequest{
		Message: message,
	}

	err := commentReplyArg.Validate()
	require.NoError(t, err)

	commentReply2, err := testRepository.CommentReply.UpdateCommentReply(commentReply1.ID, *commentReplyArg)
	require.NoError(t, err)
	require.NotEmpty(t, commentReply2)

	require.Equal(t, message, commentReply2.Message)
	require.Equal(t, commentReply1.CommentID, commentReply2.CommentID)
	require.Equal(t, commentReply1.UserID, commentReply2.UserID)

	require.NotEqual(t, commentReply1.Message, commentReply2.Message)

	// Error not found
	_, err = testRepository.CommentReply.UpdateCommentReply(0, *commentReplyArg)
	require.Error(t, err)
}

func TestDeleteCommentReply(t *testing.T) {
	commentReply1 := createRandomCommentReply(t)

	_, err := testRepository.CommentReply.DeleteCommentReply(commentReply1.ID)
	require.NoError(t, err)

	_, err = testRepository.CommentReply.GetCommentReplyById(commentReply1.ID)
	require.Error(t, err)

	// Error not found
	_, err = testRepository.CommentReply.DeleteCommentReply(0)
	require.Error(t, err)
}
