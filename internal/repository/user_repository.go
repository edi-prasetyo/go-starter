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
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByID(ctx context.Context, id int) (*model.User, error)
	UpdateVerification(ctx context.Context, email string) error
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

	query := `SELECT id, name, email, password, is_verified, otp_code, otp_expired_at 
              FROM users WHERE email = ?`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,           // id
		&user.Name,         // name
		&user.Email,        // email
		&user.Password,     // password
		&user.IsVerified,   // is_verified
		&user.OTPCode,      // otp_code
		&user.OTPExpiredAt, // otp_expired_at
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}

	query := `SELECT id, name, email, password, is_verified, otp_code, otp_expired_at 
              FROM users WHERE id = ? AND deleted_at IS NULL`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsVerified,
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
