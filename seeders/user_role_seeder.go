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

func SeedUserRoles(db *sql.DB) {
	userRoles := []struct {
		UserID int
		RoleID int
	}{
		{1, 1},
		{2, 2},
	}

	for _, ur := range userRoles {

		_, err := db.Exec(`
			INSERT INTO user_roles (user_id, role_id)
			VALUES (?, ?)
			ON DUPLICATE KEY UPDATE user_id = user_id
		`, ur.UserID, ur.RoleID)

		if err != nil {
			log.Println("failed seed user_role:", ur, err)
		}
	}

	log.Println("user_role seeder done")
}
