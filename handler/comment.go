package handler

import (
	"net/http"
	"strconv"

	"github.com/alifdwt/haiwangram/domain/requests/comment"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initCommentGroup(api *gin.Engine) {
	auth := api.Group("/api/comments")

	auth.GET("/", h.handlerGetCommentAll)
	auth.GET("/:commentId", h.handlerGetCommentById)

	auth.Use(authMiddleware(h.tokenMaker))
	auth.POST("/", h.handlerCreateComment)
	auth.PUT("/:commentId", h.handlerUpdateComment)
	auth.DELETE("/:commentId", h.handlerDeleteComment)
}

// handlerGetCommentAll function
// @Summary Get all comments
// @Description Get all comments from the application
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} []comment.CommentResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comments [get]
func (h *Handler) handlerGetCommentAll(c *gin.Context) {
	res, err := h.services.Comment.GetCommentAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerGetCommentById function
// @Summary Get comment by id
// @Description Get comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Success 200 {object} comment.CommentResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comments/{commentId} [get]
func (h *Handler) handlerGetCommentById(c *gin.Context) {
	commentIdStr := c.Param("commentId")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Comment.GetCommentById(commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerCreateComment function
// @Summary Create comment
// @Description Create comment
// @Tags comments
// @Accept json
// @Produce json
// @Param data body comment.CreateCommentRequest true "comment data"
// @Security BearerAuth
// @Success 201 {object} comment.CommentResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comments [post]
func (h *Handler) handlerCreateComment(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var body comment.CreateCommentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Comment.CreateComment(userId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerUpdateComment function
// @Summary Update comment
// @Description Update comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Param data body comment.UpdateCommentRequest true "comment data"
// @Security BearerAuth
// @Success 200 {object} comment.CommentResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comments/{commentId} [put]
func (h *Handler) handlerUpdateComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
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

	commentData, err := h.services.Comment.GetCommentById(commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if commentData.UserID != userId {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var body comment.UpdateCommentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Comment.UpdateComment(commentId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeleteComment function
// @Summary Delete comment
// @Description Delete comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Security BearerAuth
// @Success 200 {object} comment.CommentResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comments/{commentId} [delete]
func (h *Handler) handlerDeleteComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
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

	commentData, err := h.services.Comment.GetCommentById(commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if commentData.UserID != userId {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	res, err := h.services.Comment.DeleteComment(commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
