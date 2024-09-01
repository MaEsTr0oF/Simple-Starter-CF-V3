package services

import (
	"api/pkg/models"
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(dataUser models.DataUser) error {
	recipientEmail := os.Getenv("RECIPIENT_EMAIL")
	smtp_host := os.Getenv("SMTP_HOST")
	smtp_port := os.Getenv("SMTP_PORT")
	smtp_user := os.Getenv("SMTP_USER")
	smtp_password := os.Getenv("SMTP_PASSWORD")

	subject := "Subject: Я - новый участник свадьбы!\n"
	fromHeader := "From: " + smtp_user + "\n"
	toHeader := "To: " + recipientEmail +"\n"
	replyTo := "Reply-To: " + smtp_user + "\n"
	contentType := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body := fmt.Sprintf("<h1>Свадьба</h1><p>Новый участник(и): %s</p>", dataUser.FirstAndLastNames)

	msg := []byte(subject + fromHeader + toHeader + replyTo + contentType + body)

	auth := smtp.PlainAuth("", smtp_user, smtp_password, smtp_host)

	err := smtp.SendMail(smtp_host+":"+smtp_port, auth, smtp_user, []string{recipientEmail}, msg)
	if err != nil {
		return err
	}

	return nil
}
