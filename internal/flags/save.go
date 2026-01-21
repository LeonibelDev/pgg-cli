package flags

import (
	"fmt"

	"github.com/leonibeldev/pgg-cli/internal/db"
	"github.com/leonibeldev/pgg-cli/internal/models"
)

func Save(service string, user string, password string) {
	if user == "" {
		user = "default"
	}

	if service != "" {
		form := models.PasswordsForm{
			Service:  service,
			Username: user,
			Password: password,
		}

		result := db.Save(form)

		if !result {
			fmt.Printf("Error saving the password")
		}

		fmt.Println("Password saved successfully")
	}
}
