package repository

import (
	"errors"

	"github.com/alifdwt/haiwangram/domain/requests/bookmark"
	"github.com/alifdwt/haiwangram/models"
	"gorm.io/gorm"
)

type bookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) *bookmarkRepository {
	return &bookmarkRepository{db}
}

func (r *bookmarkRepository) CreateBookmark(userId int, request bookmark.CreateBookmarkRequest) (*models.Bookmark, error) {
	var bookmarkModel models.Bookmark

	db := r.db.Model(&bookmarkModel)

	bookmarkModel.UserID = userId
	bookmarkModel.PhotoID = request.PhotoID

	result := db.Debug().Create(&bookmarkModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmarkModel, nil
}

func (r *bookmarkRepository) GetBookmarkById(bookmarkId int) (*models.Bookmark, error) {
	var bookmarkModel models.Bookmark

	db := r.db.Model(&bookmarkModel)

	result := db.Debug().Preload("Photo.User").Where("id = ?", bookmarkId).First(&bookmarkModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmarkModel, nil
}

func (r *bookmarkRepository) GetBookmarkByUserId(userId int) (*[]models.Bookmark, error) {
	var bookmarkModel []models.Bookmark

	db := r.db.Model(&bookmarkModel)

	result := db.Debug().Preload("Photo.User").Where("user_id = ?", userId).Find(&bookmarkModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmarkModel, nil
}

func (r *bookmarkRepository) DeleteBookmark(photoId int, userId int) (*models.Bookmark, error) {
	var bookmarkModel models.Bookmark

	db := r.db.Model(&bookmarkModel)

	checkBookmarkById := db.Debug().Where("photo_id = ? AND user_id = ?", photoId, userId).First(&bookmarkModel)
	if checkBookmarkById.RowsAffected < 1 {
		return &bookmarkModel, errors.New("bookmark not found")
	}

	deleteBookmark := db.Debug().Delete(&bookmarkModel)
	if deleteBookmark.Error != nil {
		return &bookmarkModel, deleteBookmark.Error
	}

	return &bookmarkModel, nil
}
