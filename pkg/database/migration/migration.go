package migration

import (
	"github.com/alifdwt/haiwangram/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Photo{},
		&models.Comment{},
		&models.CommentReply{},
		&models.Like{},
		&models.SocialMedia{},
	)

	if err != nil {
		panic(err)
	}

	return err
}
