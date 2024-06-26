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
		&models.Bookmark{},
		&models.Follow{},
	)

	if err != nil {
		panic(err)
	}

	return err
}

func DropTable(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&models.User{},
		&models.Photo{},
		&models.Comment{},
		&models.CommentReply{},
		&models.Like{},
		&models.SocialMedia{},
		&models.Bookmark{},
	)

	if err != nil {
		panic(err)
	}

	return err
}
