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

func SeedRolePermissions(db *sql.DB) {

	rolePermissions := map[int][]int{

		1: {
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10,
			11, 12,
		},
	}

	for roleID, permissions := range rolePermissions {

		for _, permID := range permissions {

			_, err := db.Exec(`
				INSERT INTO role_permissions (role_id, permission_id)
				VALUES (?, ?)
				ON DUPLICATE KEY UPDATE role_id = role_id
			`, roleID, permID)

			if err != nil {
				log.Println("failed seed role_permission:", roleID, permID, err)
			}
		}
	}

	log.Println("role_permission seeder done")
}
