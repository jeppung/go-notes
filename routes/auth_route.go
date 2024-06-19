package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jeppung/go-notes.git/controllers/auth"
	repositories "github.com/jeppung/go-notes.git/repositories/auth"
	services "github.com/jeppung/go-notes.git/services/auth"
)

func Auth(r *gin.Engine) {
	authRepository := repositories.NewAuthRepository()
	authService := services.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	r.POST("/signup", authController.SignupHandler)
}
