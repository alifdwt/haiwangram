package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/models"
)

type UserMapping interface {
	ToUserResponse(request *models.User) *responses.UserResponse
	ToUserResponses(requests *[]models.User) []responses.UserResponse
}

type PhotoMapping interface {
	ToPhotoResponse(request *models.Photo) *responses.PhotoResponse
	ToPhotoWithRelationResponse(request *models.Photo) *responses.PhotoWithRelationResponse
	ToPhotoWithRelationResponses(requests *[]models.Photo) []responses.PhotoWithRelationResponse
}

type CommentMapping interface {
	ToCommentResponse(request *models.Comment) *responses.CommentResponse
	ToCommentWithRelationResponse(request *models.Comment) *responses.CommentWithRelationResponse
	ToCommentWithRelationResponses(requests *[]models.Comment) []responses.CommentWithRelationResponse
}

type CommentReplyMapping interface {
	ToCommentReplyResponse(request *models.CommentReply) *responses.CommentReplyResponse
	ToCommentReplyWithRelationResponse(request *models.CommentReply) *responses.CommentReplyWithRelationResponse
	ToCommentReplyWithRelationResponses(requests *[]models.CommentReply) []responses.CommentReplyWithRelationResponse
}

type LikeMapping interface {
	ToLikeResponse(request *models.Like) *responses.LikeResponse
	ToLikeWithRelationResponse(request *models.Like) *responses.LikeWithRelationResponse
	ToLikeWithRelationResponses(requests *[]models.Like) []responses.LikeWithRelationResponse
}
