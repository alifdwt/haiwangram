package handler

import (
	"net/http"
	"strconv"

	"github.com/alifdwt/haiwangram/domain/requests/follow"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initFollowGroup(api *gin.Engine) {
	follow := api.Group("/api/follows")

	follow.Use(authMiddleware(h.tokenMaker))
	follow.POST("/", h.handlerCreateFollow)
	follow.GET("/followed", h.handlerGetFollowedByUserId)
	follow.GET("/follower", h.handlerGetFollowerByUserId)
	follow.DELETE("/", h.handlerDeleteFollow)
}

// handlerCreateFollow function
// @Summary Create follow
// @Description Create follow on logged in user
// @Tags follows
// @Accept json
// @Produce json
// @Param data body follow.CreateFollowRequest true "follow data"
// @Security BearerAuth
// @Success 201 {object} models.Follow
// @Failure 400 {object} responses.ErrorMessage
// @Router /follows [post]
func (h *Handler) handlerCreateFollow(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var body follow.CreateFollowRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Follow.CreateFollow(userId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerGetFollowByUserId function
// @Summary Get follow by user id
// @Description Get follow by user id
// @Tags follows
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Follow
// @Failure 400 {object} responses.ErrorMessage
// @Router /follows/followed [get]
func (h *Handler) handlerGetFollowedByUserId(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	res, err := h.services.Follow.GetFollowedByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerGetFollowerByUserId function
// @Summary Get follower by user id
// @Description Get follower by user id
// @Tags follows
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Follow
// @Failure 400 {object} responses.ErrorMessage
// @Router /follows/follower [get]
func (h *Handler) handlerGetFollowerByUserId(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	res, err := h.services.Follow.GetFollowerByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeleteFollow function
// @Summary Delete follow
// @Description Delete follow
// @Tags follows
// @Accept json
// @Produce json
// @Param data body follow.DeleteFollowRequest true "follow data"
// @Security BearerAuth
// @Success 200 {object} models.Follow
// @Failure 400 {object} responses.ErrorMessage
// @Router /follows [delete]
func (h *Handler) handlerDeleteFollow(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var body follow.DeleteFollowRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Follow.DeleteFollow(userId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
