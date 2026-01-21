package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/leonibeldev/pgg-cli/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "./pgg.db")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	return conn, nil
}

func CreateDB() {
	conn, _ := Connect()

	conn.Exec(`
		CREATE TABLE IF NOT EXISTS passwords (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			service TEXT NOT NULL,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
}

func Save(form models.PasswordsForm) (result bool) {

	conn, _ := Connect()

	query := `
		INSERT INTO passwords (service, username, password)
		VALUES (?, ?, ?)
	`
	_, err := conn.Exec(query, form.Service, form.Username, form.Password)
	if err != nil {
		return false
	}

	return true
}

func GetPasswords() []models.PasswordsForm {
	conn, _ := Connect()

	var passwords []models.PasswordsForm

	query := `
		SELECT * FROM passwords
	`

	rows, err := conn.Query(query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var passwordModel models.PasswordsForm

		if err := rows.Scan(
			&passwordModel.ID,
			&passwordModel.Service,
			&passwordModel.Username,
			&passwordModel.Password,
			&passwordModel.Created_at); err != nil {
			panic(err)
		}

		passwords = append(passwords, passwordModel)
	}

	return passwords
}

func DeletePassword(ID int) {
	conn, _ := Connect()

	query := `
		DELETE FROM passwords
		WHERE id = ?
	`

	conn.Exec(query, ID)
}
