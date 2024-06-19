package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeppung/go-notes.git/common"
	"github.com/jeppung/go-notes.git/models"
	services "github.com/jeppung/go-notes.git/services/auth"
)

type AuthControllerImpl struct {
	service services.AuthService
}

func (c *AuthControllerImpl) SignupHandler(ctx *gin.Context) {
	var body models.Auth

	if err := common.JSONBindingHandler(ctx, &body); err != nil {
		return
	}

	if err := c.service.Signup(body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response[string]{
			Status: "error",
			Data:   "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response[string]{
		Status:  "success",
		Message: "User created successfully",
	})
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		service: authService,
	}
}
