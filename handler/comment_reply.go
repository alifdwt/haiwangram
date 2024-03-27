package handler

import (
	"errors"
	"net/http"
	"strconv"

	commentreply "github.com/alifdwt/haiwangram/domain/requests/comment_reply"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initCommentReplyGroup(api *gin.Engine) {
	commentReply := api.Group("/api/comment-replies")

	commentReply.GET("/", h.handlerGetCommentReplyAll)
	commentReply.GET("/:commentReplyId", h.handlerGetCommentReplyById)

	commentReply.Use(authMiddleware(h.tokenMaker))
	commentReply.POST("/", h.handlerCreateCommentReply)
	commentReply.PUT("/:commentReplyId", h.handlerUpdateCommentReply)
	commentReply.DELETE("/:commentReplyId", h.handlerDeleteCommentReply)
}

// handlerGetCommentReplyAll function
// @Summary Get all comment reply
// @Description Get all comment reply from the application
// @Tags comment reply
// @Accept json
// @Produce json
// @Success 200 {object} []responses.CommentReplyWithRelationResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comment-replies [get]
func (h *Handler) handlerGetCommentReplyAll(c *gin.Context) {
	res, err := h.services.CommentReply.GetCommentReplyAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerGetCommentReplyById function
// @Summary Get comment reply by id
// @Description Get comment reply by id
// @Tags comment reply
// @Accept json
// @Produce json
// @Param commentReplyId path int true "Comment Reply ID"
// @Success 200 {object} responses.CommentReplyWithRelationResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comment-replies/{commentReplyId} [get]
func (h *Handler) handlerGetCommentReplyById(c *gin.Context) {
	commentReplyIdStr := c.Param("commentReplyId")
	commentReplyId, err := strconv.Atoi(commentReplyIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.CommentReply.GetCommentReplyById(commentReplyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerCreateCommentReply function
// @Summary Create comment reply
// @Description Create comment reply
// @Tags comment reply
// @Accept json
// @Produce json
// @Param data body commentreply.CreateCommentReplyRequest true "comment reply data"
// @Security BearerAuth
// @Success 201 {object} responses.CommentReplyResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comment-replies [post]
func (h *Handler) handlerCreateCommentReply(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var body commentreply.CreateCommentReplyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.CommentReply.CreateCommentReply(userId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerUpdateCommentReply function
// @Summary Update comment reply
// @Description Update comment reply
// @Tags comment reply
// @Accept json
// @Produce json
// @Param commentReplyId path int true "Comment Reply ID"
// @Param data body commentreply.UpdateCommentReplyRequest true "comment reply data"
// @Security BearerAuth
// @Success 200 {object} responses.CommentReplyResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comment-replies/{commentReplyId} [put]
func (h *Handler) handlerUpdateCommentReply(c *gin.Context) {
	commentReplyIdStr := c.Param("commentReplyId")
	commentReplyId, err := strconv.Atoi(commentReplyIdStr)
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

	commentReplyData, err := h.services.CommentReply.GetCommentReplyById(commentReplyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if commentReplyData.UserID != userId {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to edit this comment reply")))
		return
	}

	var body commentreply.UpdateCommentReplyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.CommentReply.UpdateCommentReply(commentReplyId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeleteCommentReply function
// @Summary Delete comment reply
// @Description Delete comment reply
// @Tags comment reply
// @Accept json
// @Produce json
// @Param commentReplyId path int true "Comment Reply ID"
// @Security BearerAuth
// @Success 200 {object} responses.CommentReplyResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /api/comment-replies/{commentReplyId} [delete]
func (h *Handler) handlerDeleteCommentReply(c *gin.Context) {
	commentReplyIdStr := c.Param("commentReplyId")
	commentReplyId, err := strconv.Atoi(commentReplyIdStr)
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

	commentReplyData, err := h.services.CommentReply.GetCommentReplyById(commentReplyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if userId != commentReplyData.UserID {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to delete this comment reply")))
		return
	}

	res, err := h.services.CommentReply.DeleteCommentReply(commentReplyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
