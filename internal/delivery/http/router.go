// =========================================
//  Project     : Go Starter API
//  Author      : Edi Prasetyo
//  Website     : https://grahastudio.com
//  Email       : ediprasetiyo2@gmail.com
//  Version     : 1.0.0
//  License     : MIT
// =========================================
// Description:
// REST API backend using Gin, MySQL, JWT, RBAC
// =========================================

package http

import (
	"database/sql"
	"go-starter/internal/constant"
	"go-starter/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	route := gin.Default()
	rbac := middleware.NewRBAC(db)
	ctrl := NewControllerContainer(db)

	route.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain; charset=utf-8")
		c.String(200, constant.Banner)
	})

	// --- PUBLIC ROUTES ---
	auth := route.Group("/auth")
	{
		auth.POST("/register", ctrl.Auth.Register)
		auth.POST("/verify", ctrl.Auth.VerifyOTP)
		auth.POST("/login", ctrl.Auth.Login)
		auth.POST("/refresh", ctrl.Auth.Refresh)
	}

	// --- PROTECTED ROUTES ---
	auth = route.Group("/")
	auth.Use(middleware.AuthMiddleware())

	{
		auth.GET("/profile",
			rbac.RequirePermission("user.read"),
			ctrl.User.GetProfile,
		)
	}

	return route
}
