package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alifdwt/haiwangram/domain/requests/auth"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuthGroup(api *gin.Engine) {
	auth := api.Group("/api/auth")

	auth.POST("/register", h.handlerRegister)
	auth.POST("/login", h.handlerLogin)
}

// handlerRegister
// @Summary Register to the HaiwanGram
// @Description Register a new user
// @Tags auth
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "Email"
// @Param username formData string true "Username"
// @Param full_name formData string true "Full Name"
// @Param birth_date formData string true "Birth Date (YYYY-MM-DD)"
// @Param profile_image formData file true "Profile Image"
// @Param password formData string true "Password"
// @Param description formData string false "Description"
// @Success 201 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /auth/register [post]
func (h *Handler) handlerRegister(c *gin.Context) {
	email := c.Request.FormValue("email")
	username := c.Request.FormValue("username")
	fullName := c.Request.FormValue("full_name")
	birthDateString := c.Request.FormValue("birth_date")
	password := c.Request.FormValue("password")
	description := c.Request.FormValue("description")

	birthDate, err := time.Parse("2006-01-02", birthDateString)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	minimalAge := 13
	if time.Now().Year()-birthDate.Year() < minimalAge {
		c.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("user must be at least %d years old", minimalAge)))
		return
	}

	file, err := c.FormFile("profile_image")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	uploadedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer uploadedFile.Close()

	fileName := fmt.Sprintf("%s_profile", username)
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	registerReq := &auth.RegisterRequest{
		Email:           email,
		Username:        username,
		FullName:        fullName,
		BirthDate:       birthDate,
		ProfileImageURL: imageUrl,
		Password:        password,
		Description:     description,
	}

	if err := registerReq.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Auth.Register(registerReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerLogin
// @Summary Login to the HaiwanGram
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body auth.LoginRequest true "data"
// @Success 200 {object} responses.Token
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /auth/login [post]
func (h *Handler) handlerLogin(c *gin.Context) {
	var body auth.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Auth.Login(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
