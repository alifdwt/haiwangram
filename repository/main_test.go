package repository

import (
	"os"
	"testing"

	"github.com/alifdwt/haiwangram/pkg/database/migration"
	"github.com/alifdwt/haiwangram/pkg/database/postgres"
	"github.com/alifdwt/haiwangram/pkg/dotenv"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"go.uber.org/zap"
)

var testRepository *Repositories

func TestMain(m *testing.M) {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	config, err := dotenv.LoadConfig("..")
	if err != nil {
		log.Error("Error while load environtment variables", zap.Error(err))
	}

	db, err := postgres.NewClient(
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
		config.APP_TIMEZONE,
	)
	if err != nil {
		log.Error("Error while connecting to database", zap.Error(err))
	}

	err = migration.RunMigration(db)
	if err != nil {
		log.Error("Error while running migration", zap.Error(err))
	}

	repository := NewRepository(db)

	testRepository = repository

	exitCode := m.Run()

	// err = migration.DropTable(db)
	// if err != nil {
	// 	log.Error("Error while running migration", zap.Error(err))
	// }

	os.Exit(exitCode)
}
