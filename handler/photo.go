package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alifdwt/haiwangram/domain/requests/photo"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func (h *Handler) initPhotoGroup(api *gin.Engine) {
	photo := api.Group("/api/photos")

	photo.GET("/", h.handlerGetPhotoAll)
	photo.GET("/:photoId", h.handlerGetPhotoById)

	photo.Use(authMiddleware(h.tokenMaker))
	photo.POST("/", h.handlerCreatePhoto)
	photo.PUT("/:photoId", h.handlerUpdatePhoto)
	photo.DELETE("/:photoId", h.handlerDeletePhoto)
}

// handlerGetPhotoAll function
// @Summary Get all photos
// @Description Get all photos from the application
// @Tags photos
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Success 200 {object} []responses.PhotoWithRelationResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos [get]
func (h *Handler) handlerGetPhotoAll(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Photo.GetPhotoAll(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerGetPhotoById function
// @Summary Get photo by id
// @Description Get photo by id
// @Tags photos
// @Accept json
// @Produce json
// @Param photoId path int true "Photo ID"
// @Success 200 {object} responses.PhotoWithRelationResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos/{photoId} [get]
func (h *Handler) handlerGetPhotoById(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Photo.GetPhotoById(photoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerCreatePhoto function
// @Summary Create photo
// @Description Create photo
// @Tags photos
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Title"
// @Param caption formData string true "Caption"
// @Param image formData file true "Image"
// @Success 201 {object} responses.PhotoResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos [post]
func (h *Handler) handlerCreatePhoto(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	title := c.Request.FormValue("title")
	caption := c.Request.FormValue("caption")

	imageFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	uploadedFile, err := imageFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer uploadedFile.Close()

	fileName := fmt.Sprintf("%s_%s", authPayload.Username, slug.Make(title))
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	createPhotoReq := &photo.CreatePhotoRequest{
		Title:    title,
		Caption:  caption,
		PhotoURL: imageUrl,
	}

	if err := createPhotoReq.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Photo.CreatePhoto(userId, *createPhotoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerUpdatePhoto function
// @Summary Update photo
// @Description Update photo
// @Tags photos
// @Accept multipart/form-data
// @Produce json
// @Param photoId path int true "Photo ID"
// @Param title formData string true "Title"
// @Param caption formData string true "Caption"
// @Param image formData file true "Image"
// @Security BearerAuth
// @Success 200 {object} responses.PhotoResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos/{photoId} [put]
func (h *Handler) handlerUpdatePhoto(c *gin.Context) {
	photoIdStr := c.Param("photoId")
	photoId, err := strconv.Atoi(photoIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	photoData, err := h.services.Photo.GetPhotoById(photoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if photoData.UserID != userId {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to update this photo")))
		return
	}

	title := c.Request.FormValue("title")
	caption := c.Request.FormValue("caption")

	imageFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	uploadedFile, err := imageFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer uploadedFile.Close()

	fileName := fmt.Sprintf("%s_%s", authPayload.Username, slug.Make(title))
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	updatePhotoReq := &photo.UpdatePhotoRequest{
		Title:    title,
		Caption:  caption,
		PhotoURL: imageUrl,
	}

	if err := updatePhotoReq.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Photo.UpdatePhoto(userId, photoId, *updatePhotoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeletePhoto function
// @Summary Delete photo
// @Description Delete photo
// @Tags photos
// @Accept json
// @Produce json
// @Param photoId path int true "Photo ID"
// @Security BearerAuth
// @Success 200 {object} responses.PhotoResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos/{photoId} [delete]
func (h *Handler) handlerDeletePhoto(c *gin.Context) {
	photoIdStr := c.Param("photoId")
	photoId, err := strconv.Atoi(photoIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	photoData, err := h.services.Photo.GetPhotoById(photoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if photoData.UserID != userId {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to delete this photo")))
		return
	}

	res, err := h.services.Photo.DeletePhoto(photoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
