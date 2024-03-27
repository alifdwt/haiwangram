package handler

import (
	"net/http"
	"strconv"

	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserGroup(api *gin.Engine) {
	user := api.Group("/api/users")

	user.Use(authMiddleware(h.tokenMaker))
	user.DELETE("/", h.handlerDeleteUser)
}

// handlerDeleteUser function
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} user.UserResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /users [delete]
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
