package repository

import (
	"errors"

	"github.com/alifdwt/haiwangram/domain/requests/photo"
	"github.com/alifdwt/haiwangram/models"
	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) CreatePhoto(userId int, request photo.CreatePhotoRequest) (*models.Photo, error) {
	var photoModel models.Photo

	db := r.db.Model(&photoModel)

	photoModel.Title = request.Title
	photoModel.Caption = request.Caption
	photoModel.PhotoURL = request.PhotoURL
	photoModel.UserID = userId

	result := db.Debug().Create(&photoModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &photoModel, nil
}

func (r *photoRepository) GetPhotoAll(limit int) (*[]models.Photo, error) {
	var photos []models.Photo

	db := r.db.Model(&photos)

	result := db.Debug().Preload("User").Preload("Comments").Preload("Likes").Preload("Bookmarks").Order("created_at DESC").Limit(limit).Find(&photos)
	if result.Error != nil {
		return nil, result.Error
	}

	return &photos, nil
}

func (r *photoRepository) GetPhotoById(photoId int) (*models.Photo, error) {
	var photo models.Photo

	db := r.db.Model(&photo)

	result := db.Debug().Preload("User").Preload("Comments").Preload("Likes").Preload("Bookmarks").Where("id = ?", photoId).First(&photo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &photo, nil
}

func (r *photoRepository) GetPhotoByUserId(userId int, limit int) (*[]models.Photo, error) {
	var photos []models.Photo

	db := r.db.Model(&photos)

	result := db.Debug().Preload("User").Preload("Comments").Preload("Likes").Preload("Bookmarks").Where("user_id = ?", userId).Order("created_at DESC").Limit(limit).Find(&photos)
	if result.Error != nil {
		return nil, result.Error
	}

	return &photos, nil
}

func (r *photoRepository) UpdatePhoto(userId int, photoId int, updatedPhoto photo.UpdatePhotoRequest) (*models.Photo, error) {
	var photo models.Photo

	db := r.db.Model(&photo)

	checkPhotoById := db.Debug().Where("id = ? AND user_id = ?", photoId, userId).First(&photo)
	if checkPhotoById.RowsAffected > 1 {
		return &photo, errors.New("photo not found")
	}

	photo.Title = updatedPhoto.Title
	photo.PhotoURL = updatedPhoto.PhotoURL
	photo.Caption = updatedPhoto.Caption

	updatePhoto := db.Debug().Updates(&photo)
	if updatePhoto.Error != nil {
		return &photo, updatePhoto.Error
	}

	return &photo, nil
}

func (r *photoRepository) DeletePhoto(photoId int) (*models.Photo, error) {
	var photo models.Photo

	db := r.db.Model(&photo)

	checkPhotoById := db.Debug().Where("id = ?", photoId).First(&photo)
	if checkPhotoById.RowsAffected < 1 {
		return &photo, errors.New("photo not found")
	}

	deletePhoto := db.Debug().Delete(&photo)
	if deletePhoto.Error != nil {
		return &photo, deletePhoto.Error
	}

	return &photo, nil
}
