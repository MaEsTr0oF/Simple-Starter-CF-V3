package services

import (
	"api/pkg/models"
	"net/smtp"
	"os"
)

func SendEmail(dataUser models.DataUser) error {
	recipientEmail := os.Getenv("RECIPIENT_EMAIL")
	smtp_host := os.Getenv("SMTP_HOST")
	smtp_port := os.Getenv("SMTP_PORT")
	smtp_user := os.Getenv("SMTP_USER")
	smtp_password := os.Getenv("SMTP_PASSWORD")

	msg := []byte("Subject: Проверка\r\n\r\n" + dataUser.FirstAndLastNames)

	auth := smtp.PlainAuth("", smtp_user, smtp_password, smtp_host)

	err := smtp.SendMail(smtp_host+":"+smtp_port, auth, smtp_user, []string{recipientEmail}, msg)
	if err != nil {
		return err
	}

	return nil
}
