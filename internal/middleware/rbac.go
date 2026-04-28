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

package middleware

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RBAC struct {
	DB *sql.DB
}

func NewRBAC(db *sql.DB) *RBAC {
	return &RBAC{DB: db}
}

func (r *RBAC) RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, ok := c.Get("user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		var exists int
		err := r.DB.QueryRow(`
			SELECT 1
			FROM user_roles ur
			JOIN role_permissions rp ON rp.role_id = ur.role_id
			JOIN permissions p ON p.id = rp.permission_id
			WHERE ur.user_id = ? AND p.name = ?
			LIMIT 1
		`, userID, permission).Scan(&exists)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
