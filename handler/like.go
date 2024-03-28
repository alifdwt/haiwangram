package handler

import (
	"net/http"
	"strconv"

	"github.com/alifdwt/haiwangram/domain/requests/like"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initLikeGroup(api *gin.Engine) {
	like := api.Group("/api/likes")

	like.Use(authMiddleware(h.tokenMaker))
	like.POST("/", h.handlerCreateLike)
	like.DELETE("/:likeId", h.handlerDeleteLike)
}

// handlerCreateLike function
// @Summary Create like
// @Description Create like
// @Tags likes
// @Accept json
// @Produce json
// @Param data body like.CreateLikeRequest true "like data"
// @Security BearerAuth
// @Success 201 {object} responses.LikeResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /likes [post]
func (h *Handler) handlerCreateLike(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var body like.CreateLikeRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Like.CreateLike(userId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerDeleteLike function
// @Summary Delete like
// @Description Delete like by id
// @Tags likes
// @Accept json
// @Produce json
// @Param likeId path int true "Like ID"
// @Security BearerAuth
// @Success 200 {object} responses.LikeResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /likes/{likeId} [delete]
func (h *Handler) handlerDeleteLike(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	likeIdStr := c.Param("likeId")
	likeId, err := strconv.Atoi(likeIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	likeData, err := h.services.Like.GetLikeById(likeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if likeData.UserID != userId {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	res, err := h.services.Like.DeleteLike(likeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
