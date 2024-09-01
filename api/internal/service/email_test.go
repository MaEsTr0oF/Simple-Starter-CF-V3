package services

import (
	"api/pkg/models"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	assert.NotEqual(t, "", os.Getenv("SMTP_HOST"))

	var dataTest = models.DataUser{
		FirstAndLastNames: "Адель Альзоаби",
		Attendance: 1,
	}

	err := SendEmail(dataTest)
	assert.Nil(t, err)
}
