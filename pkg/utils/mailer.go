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

package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendOTPEmail(to string, otp string) error {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	message := []byte(fmt.Sprintf("Subject: Your OTP Code\r\n\r\nYour OTP code is: %s. It will expire in 5 minutes.", otp))

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)

	return err
}
