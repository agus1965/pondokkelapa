package main

import (
	"fmt"

	"github.com/agusrio/boxith/config"
	"github.com/agusrio/boxith/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	authController controllers.AuthController = controllers.NewAuthContoller()
)

func main() {
	fmt.Println("Ones")
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}
