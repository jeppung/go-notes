package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	SignupHandler(ctx *gin.Context)
}
