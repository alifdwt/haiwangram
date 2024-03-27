package repository

import (
	"errors"

	commentreply "github.com/alifdwt/haiwangram/domain/requests/comment_reply"
	"github.com/alifdwt/haiwangram/models"
	"gorm.io/gorm"
)

type commentReplyRepository struct {
	db *gorm.DB
}

func NewCommentReplyRepository(db *gorm.DB) *commentReplyRepository {
	return &commentReplyRepository{db}
}

func (r *commentReplyRepository) CreateCommentReply(userId int, request commentreply.CreateCommentReplyRequest) (*models.CommentReply, error) {
	var commentReplyModel models.CommentReply

	db := r.db.Model(&commentReplyModel)

	commentReplyModel.Message = request.Message
	commentReplyModel.CommentID = request.CommentID
	commentReplyModel.UserID = userId

	result := db.Debug().Create(&commentReplyModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &commentReplyModel, nil
}

func (r *commentReplyRepository) GetCommentReplyAll() (*[]models.CommentReply, error) {
	var commentReplies []models.CommentReply

	db := r.db.Model(&commentReplies)

	result := db.Debug().Preload("User").Preload("Comment").Find(&commentReplies)
	if result.Error != nil {
		return nil, result.Error
	}

	return &commentReplies, nil
}

func (r *commentReplyRepository) GetCommentReplyById(commentReplyId int) (*models.CommentReply, error) {
	var commentReply models.CommentReply

	db := r.db.Model(&commentReply)

	result := db.Debug().Preload("User").Preload("Comment").Where("id = ?", commentReplyId).First(&commentReply)
	if result.Error != nil {
		return nil, result.Error
	}

	return &commentReply, nil
}

func (r *commentReplyRepository) UpdateCommentReply(commentReplyId int, updatedCommentReply commentreply.UpdateCommentReplyRequest) (*models.CommentReply, error) {
	var commentReply models.CommentReply

	db := r.db.Model(&commentReply)

	checkCommentReplyById := db.Debug().Where("id = ?", commentReplyId).First(&commentReply)
	if checkCommentReplyById.RowsAffected < 1 {
		return &commentReply, errors.New("comment reply not found")
	}

	commentReply.Message = updatedCommentReply.Message

	updateCommentReply := db.Debug().Updates(&commentReply)
	if updateCommentReply.Error != nil {
		return &commentReply, updateCommentReply.Error
	}

	return &commentReply, nil
}

func (r *commentReplyRepository) DeleteCommentReply(commentReplyId int) (*models.CommentReply, error) {
	var commentReply models.CommentReply

	db := r.db.Model(&commentReply)

	checkCommentReplyById := db.Debug().Where("id = ?", commentReplyId).First(&commentReply)
	if checkCommentReplyById.RowsAffected < 1 {
		return &commentReply, errors.New("comment reply not found")
	}

	deleteCommentReply := db.Debug().Delete(&commentReply)
	if deleteCommentReply.Error != nil {
		return &commentReply, deleteCommentReply.Error
	}

	return &commentReply, nil
}
