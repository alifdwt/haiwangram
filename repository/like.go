package repository

import (
	"errors"

	"github.com/alifdwt/haiwangram/domain/requests/like"
	"github.com/alifdwt/haiwangram/models"
	"gorm.io/gorm"
)

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *likeRepository {
	return &likeRepository{db}
}

func (r *likeRepository) CreateLike(userId int, request like.CreateLikeRequest) (*models.Like, error) {
	var likeModel models.Like

	db := r.db.Model(&likeModel)

	likeModel.UserID = userId
	likeModel.PhotoID = request.PhotoID

	result := db.Debug().Create(&likeModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &likeModel, nil
}

// func (r *likeRepository) GetLikeAll() (*[]models.Like, error) {
// 	var likes []models.Like

// 	db := r.db.Model(&likes)

// 	result := db.Debug().Preload("User").Preload("Photo").Find(&likes)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &likes, nil
// }

func (r *likeRepository) GetLikeById(likeId int) (*models.Like, error) {
	var like models.Like

	db := r.db.Model(&like)

	result := db.Debug().Preload("User").Preload("Photo").Where("id = ?", likeId).First(&like)
	if result.Error != nil {
		return nil, result.Error
	}

	return &like, nil
}

func (r *likeRepository) DeleteLike(photoId int, userId int) (*models.Like, error) {
	var like models.Like

	db := r.db.Model(&like)

	checkLikeById := db.Debug().Where("photo_id = ? AND user_id = ?", photoId, userId).First(&like)
	if checkLikeById.RowsAffected < 1 {
		return &like, errors.New("like not found")
	}

	deleteLike := db.Debug().Delete(&like)
	if deleteLike.Error != nil {
		return &like, deleteLike.Error
	}

	return &like, nil
}
