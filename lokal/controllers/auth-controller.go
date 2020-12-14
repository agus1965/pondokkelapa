package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController interface...
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// disini masukkan service yg kalian butuh
	// this is where you put your service
}

//NewAuthContoller creates a new instance of AuthController
func NewAuthContoller() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Login",
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Register",
	})
}
