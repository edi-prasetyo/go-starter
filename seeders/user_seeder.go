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

package seeders

import (
	"database/sql"
	"log"

	"go-starter/pkg/utils"
)

func SeedUsers(db *sql.DB) {

	type User struct {
		Name     string
		Email    string
		Password string
	}

	rawUsers := []User{
		{"Admin", "admin@mail.com", "12345678"},
		{"Warehouse", "warehouse@mail.com", "12345678"},
	}

	for _, u := range rawUsers {

		hash, err := utils.HashPassword(u.Password)
		if err != nil {
			log.Println("hash error:", err)
			continue
		}

		_, err = db.Exec(`
			INSERT INTO users (name, email, password)
			VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE email = email
		`, u.Name, u.Email, hash)

		if err != nil {
			log.Println("seed user error:", u.Email, err)
		}
	}

	log.Println("user seeder done")
}
