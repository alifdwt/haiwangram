package handler

import (
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/pkg/cloudinary"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/alifdwt/haiwangram/service"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	_ "github.com/alifdwt/haiwangram/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services   *service.Service
	cloudinary cloudinary.MyCloudinary
	tokenMaker token.Maker
}

func NewHandler(services *service.Service, cloudinary cloudinary.MyCloudinary, tokenMaker token.Maker) *Handler {
	return &Handler{
		services:   services,
		cloudinary: cloudinary,
		tokenMaker: tokenMaker,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	h.InitApi(router)

	return router
}

func (h *Handler) InitApi(router *gin.Engine) {
	h.initAuthGroup(router)
	h.initUserGroup(router)
	h.initPhotoGroup(router)
	h.initCommentGroup(router)
	h.initCommentReplyGroup(router)
	h.initLikeGroup(router)
	h.initBookmarkGroup(router)
	h.initFollowGroup(router)
}

func errorResponse(err error) responses.ErrorMessage {
	return responses.ErrorMessage{
		Message: err.Error(),
	}
}
