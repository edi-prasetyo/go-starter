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

package usecase

import (
	"context"
	"database/sql"
	"errors"
	"go-starter/internal/model"
	"go-starter/internal/repository"
	"go-starter/pkg/utils"
	"time"
)

type AuthUsecase interface {
	Register(ctx context.Context, req model.RegisterRequest) error
	VerifyOTP(ctx context.Context, req model.VerifyOTPRequest) (*utils.TokenResponse, error)
	Login(ctx context.Context, email, password string) (*utils.TokenResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*utils.TokenResponse, error)
}

type authUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return &authUsecase{repo}
}

func (u *authUsecase) Register(ctx context.Context, req model.RegisterRequest) error {
	existing, _ := u.repo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return errors.New("email sudah terdaftar")
	}

	hashedPassword, _ := utils.HashPassword(req.Password)
	otp := utils.GenerateOTP()
	expiry := time.Now().Add(5 * time.Minute)

	user := &model.User{
		Name:         req.Name,
		Email:        req.Email,
		Password:     hashedPassword,
		OTPCode:      sql.NullString{String: otp, Valid: true},
		OTPExpiredAt: &expiry,
	}

	err := u.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	go utils.SendOTPEmail(user.Email, otp)
	return nil
}

func (u *authUsecase) VerifyOTP(ctx context.Context, req model.VerifyOTPRequest) (*utils.TokenResponse, error) {
	user, err := u.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if user.OTPCode.String != req.OTP {
		return nil, errors.New("kode OTP salah")
	}

	if user.OTPExpiredAt != nil && time.Now().After(*user.OTPExpiredAt) {
		return nil, errors.New("kode OTP sudah kedaluwarsa")
	}

	err = u.repo.UpdateVerification(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	return utils.GenerateTokenResponse(user.ID, user.Email)
}

func (u *authUsecase) Login(ctx context.Context, email, password string) (*utils.TokenResponse, error) {
	// 1. Cari user dulu
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		// Jika error (misal: user tidak ketemu), langsung stop di sini
		return nil, errors.New("email atau password salah")
	}

	// 2. Sekarang aman untuk cek password karena 'user' dipastikan tidak nil
	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("email atau password salah")
	}

	if !user.IsVerified {
		return nil, errors.New("akun belum diverifikasi")
	}

	return utils.GenerateTokenResponse(user.ID, user.Email)
}

func (u *authUsecase) RefreshToken(ctx context.Context, rt string) (*utils.TokenResponse, error) {
	claims, err := utils.ValidateToken(rt)
	if err != nil || claims["type"] != "refresh" {
		return nil, errors.New("refresh token tidak valid")
	}

	userID := int(claims["user_id"].(float64))
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	return utils.GenerateTokenResponse(user.ID, user.Email)
}
