package handler

import (
	"net/http"
	"strconv"

	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserGroup(api *gin.Engine) {
	user := api.Group("/users")

	user.Use(authMiddleware(h.tokenMaker))
	user.DELETE("/", h.handlerDeleteUser)
}

func (h *Handler) handlerDeleteUser(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	res, err := h.services.User.DeleteUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
