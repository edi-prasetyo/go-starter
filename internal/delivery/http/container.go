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
	"go-starter/internal/delivery/handler"
	"go-starter/internal/repository"
	"go-starter/internal/usecase"
)

type ControllerContainer struct {
	Auth *handler.AuthHandler
	User *handler.UserHandler
}

func NewControllerContainer(db *sql.DB) *ControllerContainer {
	// Repositories
	userRepo := repository.NewUserRepository(db)

	// Usecases
	authUC := usecase.NewAuthUsecase(userRepo)
	userUC := usecase.NewUserUsecase(userRepo)

	// Controllers / Handlers
	return &ControllerContainer{
		Auth: handler.NewAuthHandler(authUC),
		User: handler.NewUserHandler(userUC),
	}
}
