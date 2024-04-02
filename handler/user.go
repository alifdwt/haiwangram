package handler

import (
	"net/http"
	"strconv"

	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserGroup(api *gin.Engine) {
	user := api.Group("/api/users")

	user.GET(("/random/:count"), h.handlerGetRandomUser)
	user.GET("/:username", h.handlerGetUserByUsername)

	user.Use(authMiddleware(h.tokenMaker))
	user.GET("/", h.handleGetLoggedInUser)
	user.DELETE("/", h.handlerDeleteUser)
}

// handlerGetLoggedInUser function
// @Summary Get logged in user
// @Description Get logged in user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /users [get]
func (h *Handler) handleGetLoggedInUser(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	email, err := h.services.User.GetUserByEmail(authPayload.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res, err := h.services.User.GetUserByEmail(email.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeleteUser function
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.UserResponse
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

// handlerGetRandomUser function
// @Summary Get random user
// @Description Get random user
// @Tags users
// @Accept json
// @Produce json
// @Param count path int true "Count"
// @Success 200 {object} []responses.UserResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /users/random/{count} [get]
func (h *Handler) handlerGetRandomUser(c *gin.Context) {
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.User.GetRandomUser(count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerGetUserByUsername function
// @Summary Get user by username
// @Description Get user by username
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /users/{username} [get]
func (h *Handler) handlerGetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	res, err := h.services.User.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
