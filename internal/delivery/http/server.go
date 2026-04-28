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
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartServer(db *sql.DB) {
	router := SetupRouter(db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	log.Printf("Server running on port %s", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen error: %s\n", err)
	}
}
