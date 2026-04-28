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

package repository

import (
	"context"
	"database/sql"
	"go-starter/internal/model"
	"log"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByID(ctx context.Context, id int) (*model.User, error)
	UpdateVerification(ctx context.Context, email string) error
	GetProfile(ctx context.Context, id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email, password, otp_code, otp_expired_at) 
              VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
		user.OTPCode.String,
		user.OTPExpiredAt,
	)
	return err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}

	query := `SELECT id, name, email, password, is_verified, fcm_token, otp_code, otp_expired_at 
              FROM users WHERE email = ?`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,         // id
		&user.Name,       // name
		&user.Email,      // email
		&user.Password,   // password
		&user.IsVerified, // is_verified
		&user.FCMToken,
		&user.OTPCode,      // otp_code
		&user.OTPExpiredAt, // otp_expired_at
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}

	query := `SELECT id, name, email, password, is_verified, fcm_token, otp_code, otp_expired_at 
              FROM users WHERE id = ? AND deleted_at IS NULL`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsVerified,
		&user.FCMToken,
		&user.OTPCode,
		&user.OTPExpiredAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdateVerification(ctx context.Context, email string) error {
	query := "UPDATE users SET is_verified = true, otp_code = NULL, otp_expiry = NULL WHERE email = ?"
	_, err := r.db.ExecContext(ctx, query, email)
	return err
}

func (r *userRepository) GetProfile(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	user.Roles = []model.Role{} // Inisialisasi agar tidak null

	log.Printf("DEBUG: Memulai GetProfile untuk ID: %d", id)

	// 1. Ambil data User
	// Coba hapus dulu created_at & updated_at dari query jika masih error "0001-01-01"
	queryUser := `SELECT id, name, email, is_verified, fcm_token, created_at, updated_at FROM users WHERE id = ?`
	err := r.db.QueryRowContext(ctx, queryUser, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.IsVerified, &user.FCMToken, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		log.Printf("DEBUG ERROR di Scan User: %v", err) // CEK LOG INI!
		return nil, err
	}

	log.Printf("DEBUG: User %s ditemukan, lanjut cari roles...", user.Name)

	// 2. Ambil data Roles
	queryRoles := `
        SELECT r.id, r.name, r.description 
        FROM roles r
        JOIN user_roles ur ON r.id = ur.role_id
        WHERE ur.user_id = ?`

	rows, err := r.db.QueryContext(ctx, queryRoles, id)
	if err != nil {
		log.Printf("DEBUG ERROR di Query Roles: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role model.Role
		var desc sql.NullString

		if err := rows.Scan(&role.ID, &role.Name, &desc); err != nil {
			log.Printf("DEBUG ERROR di Scan Role: %v", err)
			return nil, err
		}
		role.Description = desc.String

		log.Printf("DEBUG: Berhasil menambahkan Role: %s", role.Name)
		user.Roles = append(user.Roles, role)
	}

	return user, nil
}
