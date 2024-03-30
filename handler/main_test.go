package handler

import (
	"os"
	"testing"

	"github.com/alifdwt/haiwangram/pkg/cloudinary"
	"github.com/alifdwt/haiwangram/pkg/dotenv"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/alifdwt/haiwangram/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestHandler(t *testing.T, authService service.AuthService) *Handler {
	config, err := dotenv.LoadConfig("..")
	require.NoError(t, err)

	myCloudinary, err := cloudinary.NewMyCloudinary(
		config.CLOUDINARY_CLOUD_NAME,
		config.CLOUDINARY_API_KEY,
		config.CLOUDINARY_SECRET_KEY,
		config.CLOUDINARY_UPLOAD_FOLDER,
	)
	require.NoError(t, err)

	tokenMaker, err := token.NewJWTMaker(config.TOKEN_SYMETRIC_KEY)
	require.NoError(t, err)

	testService := service.Service{
		Auth: authService,
	}

	handler := NewHandler(&testService, *myCloudinary, tokenMaker)

	return handler
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
