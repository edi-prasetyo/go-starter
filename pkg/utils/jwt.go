// =========================================
//
//	Project     : Go Starter API
//	Author      : Edi Prasetyo
//	Website     : https://grahastudio.com
//	Email       : ediprasetiyo2@gmail.com
//	Version     : 1.0.0
//	License     : MIT
//
// =========================================
// Description:
// REST API backend using Gin, MySQL, JWT, RBAC
// =========================================

package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateTokenResponse(userID int, email string) (*TokenResponse, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	atClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"type":    "access",
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	atString, err := at.SignedString(secret)
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{
		"user_id": userID,
		"type":    "refresh",
		"exp":     time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rtString, err := rt.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  atString,
		RefreshToken: rtString,
	}, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
