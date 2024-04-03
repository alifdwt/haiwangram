package handler

import (
	"net/http"
	"strconv"

	"github.com/alifdwt/haiwangram/domain/requests/bookmark"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initBookmarkGroup(api *gin.Engine) {
	bookmark := api.Group("/api/bookmarks")

	bookmark.Use(authMiddleware(h.tokenMaker))
	bookmark.POST("/", h.handlerCreateBookmark)
	bookmark.GET("/", h.handlerGetBookmarkByUserId)
	bookmark.DELETE("/", h.handlerDeleteBookmark)
}

// handlerCreateBookmark function
// @Summary Create bookmark
// @Description Create bookmark
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param data body bookmark.CreateBookmarkRequest true "bookmark data"
// @Security BearerAuth
// @Success 201 {object} responses.BookmarkResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /bookmarks [post]
func (h *Handler) handlerCreateBookmark(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var body bookmark.CreateBookmarkRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Bookmark.CreateBookmark(userId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerGetBookmarkByUserId function
// @Summary Get user bookmarks
// @Description Get logged in user bookmarks
// @Tags bookmarks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.BookmarkResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /bookmarks [get]
func (h *Handler) handlerGetBookmarkByUserId(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	res, err := h.services.Bookmark.GetBookmarkByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeleteBookmark function
// @Summary Delete bookmark
// @Description Delete bookmark
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param data body bookmark.DeleteBookmarkRequest true "bookmark data"
// @Security BearerAuth
// @Success 200 {object} responses.BookmarkResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /bookmarks [delete]
func (h *Handler) handlerDeleteBookmark(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var body bookmark.DeleteBookmarkRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Bookmark.DeleteBookmark(body.PhotoID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
