package service

import (
	"os"
	"testing"

	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/database/migration"
	"github.com/alifdwt/haiwangram/pkg/database/postgres"
	"github.com/alifdwt/haiwangram/pkg/dotenv"
	"github.com/alifdwt/haiwangram/pkg/hashing"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/alifdwt/haiwangram/repository"
	"go.uber.org/zap"
)

var testService *Service

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

	hashing := hashing.NewHashingPassword()
	repository := repository.NewRepository(db)

	tokenMaker, err := token.NewJWTMaker(config.TOKEN_SYMETRIC_KEY)
	if err != nil {
		log.Error("Error while creating token maker", zap.Error(err))
	}

	mapper := mapper.NewMapper()

	service := NewService(Deps{
		Config:     config,
		Repository: repository,
		Hashing:    *hashing,
		TokenMaker: tokenMaker,
		Logger:     *log,
		Mapper:     *mapper,
	})

	testService = service

	exitCode := m.Run()

	// err = migration.DropTable(db)
	// if err != nil {
	// 	log.Error("Error while running migration", zap.Error(err))
	// }

	os.Exit(exitCode)
}
