package models

import "time"

type PasswordsForm struct {
	ID         int
	Service    string
	Username   string
	Password   string
	Created_at time.Time
}
