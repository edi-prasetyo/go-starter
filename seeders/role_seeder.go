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
)

type Role struct {
	Name        string
	Description string
}

func SeedRoles(db *sql.DB) {

	roles := []Role{
		{"administrator", "Full access to all system features"},
		{"admin", "Limited operational access"},
		{"staff", "Limited operational access"},
		{"finance", "Handle financial data"},
	}

	for _, r := range roles {

		_, err := db.Exec(`
			INSERT INTO roles (name, description)
			VALUES (?, ?)
			ON DUPLICATE KEY UPDATE name = name
		`, r.Name, r.Description)

		if err != nil {
			log.Println("failed seed role:", r.Name, err)
		}
	}

	log.Println("role seeder done")
}
