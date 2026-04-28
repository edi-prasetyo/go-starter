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

package main

import (
	"log"

	"github.com/joho/godotenv"

	"go-starter/pkg/database"
	"go-starter/seeders"
)

func main() {

	// Load env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env tidak ditemukan")
	}

	db := database.InitMySQL()

	log.Println("running seeders...")

	seeders.Run(db)

	log.Println("seeding completed")
}
