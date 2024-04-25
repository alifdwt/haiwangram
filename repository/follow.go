package repository

import (
	"errors"

	"github.com/alifdwt/haiwangram/domain/requests/follow"
	"github.com/alifdwt/haiwangram/models"
	"gorm.io/gorm"
)

type followRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) *followRepository {
	return &followRepository{db}
}

func (r *followRepository) CreateFollow(userId int, request follow.CreateFollowRequest) (*models.Follow, error) {
	var followModel models.Follow

	db := r.db.Model(&followModel)

	followModel.FollowerID = userId
	followModel.FollowedID = request.UserId

	result := db.Debug().Create(&followModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &followModel, nil
}

func (r *followRepository) GetFollowedByUserId(userId int) (*[]models.Follow, error) {
	var followModel []models.Follow

	db := r.db.Model(&followModel)

	result := db.Debug().Preload("Follower").Preload("Followed").Where("follower_id = ?", userId).Find(&followModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &followModel, nil
}

func (r *followRepository) GetFollowerByUserId(userId int) (*[]models.Follow, error) {
	var followModel []models.Follow

	db := r.db.Model(&followModel)

	result := db.Debug().Preload("Follower").Preload("Followed").Where("followed_id = ?", userId).Find(&followModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &followModel, nil
}

func (r *followRepository) DeleteFollow(userId int, request follow.DeleteFollowRequest) (*models.Follow, error) {
	var followModel models.Follow

	db := r.db.Model(&followModel)

	checkFollowById := db.Debug().Where("follower_id = ? AND followed_id = ?", userId, request.UserId).First(&followModel)
	if checkFollowById.RowsAffected < 1 {
		return &followModel, errors.New("follow not found")
	}

	deleteFollow := db.Debug().Delete(&followModel)
	if deleteFollow.Error != nil {
		return &followModel, deleteFollow.Error
	}

	return &followModel, nil
}
