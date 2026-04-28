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

import "database/sql"

func Run(db *sql.DB) {
	SeedUsers(db)
	SeedRoles(db)
	SeedPermissions(db)

	SeedUserRoles(db)
	SeedRolePermissions(db)
}
