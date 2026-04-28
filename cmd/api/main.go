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
	"os"

	"github.com/joho/godotenv"

	"go-starter/internal/banner"
	"go-starter/internal/delivery/http"
	"go-starter/pkg/database"
	"go-starter/seeders"
)

func main() {

	banner.Print()

	// load env
	godotenv.Load(".env")

	db := database.InitMySQL()
	defer db.Close()

	// migration
	database.RunMigrations(db)

	// check argument
	if len(os.Args) > 1 && os.Args[1] == "--seed" {
		log.Println("running seeders...")
		seeders.Run(db)
		log.Println("seeding completed")
		return
	}

	http.StartServer(db)
}
