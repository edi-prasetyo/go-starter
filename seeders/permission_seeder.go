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

func SeedPermissions(db *sql.DB) {

	type Permission struct {
		Name   string
		Module string
	}

	data := []Permission{
		{"user.read", "user"},
		{"user.create", "user"},
		{"user.update", "user"},
		{"user.delete", "user"},

		{"role.read", "role"},
		{"role.create", "role"},
		{"role.update", "role"},
		{"role.delete", "role"},

		{"permission.read", "permission"},
		{"permission.manage", "permission"},
	}

	for _, p := range data {
		_, err := db.Exec(`
			INSERT INTO permissions (name, module)
			VALUES (?, ?)
			ON DUPLICATE KEY UPDATE name = name
		`, p.Name, p.Module)

		if err != nil {
			log.Println("failed seed permission:", p.Name, err)
		}
	}

	log.Println("permission seeder done")
}
