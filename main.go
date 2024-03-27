package main

import (
	"github.com/alifdwt/haiwangram/handler"
	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/cloudinary"
	"github.com/alifdwt/haiwangram/pkg/database/migration"
	"github.com/alifdwt/haiwangram/pkg/database/postgres"
	"github.com/alifdwt/haiwangram/pkg/dotenv"
	"github.com/alifdwt/haiwangram/pkg/hashing"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/alifdwt/haiwangram/repository"
	"github.com/alifdwt/haiwangram/service"
	"go.uber.org/zap"
)

// @title HaiwanGram API
// @version 1.0
// @description This is HaiwanGram API documentation.
// @termsOfService http://swagger.io/terms/

// @contact.name Alif Dewantara
// @contact.url http://github.com/alifdwt
// @contact.email aputradewantara@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	config, err := dotenv.LoadConfig(".")
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

	myCloudinary, err := cloudinary.NewMyCloudinary(
		config.CLOUDINARY_CLOUD_NAME,
		config.CLOUDINARY_API_KEY,
		config.CLOUDINARY_SECRET_KEY,
		config.CLOUDINARY_UPLOAD_FOLDER,
	)
	if err != nil {
		log.Error("Error while connecting to cloudinary", zap.Error(err))
	}

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Config:     config,
		Repository: repository,
		Hashing:    *hashing,
		TokenMaker: tokenMaker,
		Logger:     *log,
		Mapper:     *mapper,
	})

	myHandler := handler.NewHandler(service, *myCloudinary, tokenMaker)

	err = myHandler.Init().Run(config.SERVER_ADDRESS)
	if err != nil {
		log.Error("Cannot start server", zap.Error(err))
	}
}
