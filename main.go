/*
Copyright © 2026 leonibeldev <leonibel.ramirez@gmail.com>
*/
package main

import (
	"github.com/leonibeldev/pgg-cli/cmd"
	"github.com/leonibeldev/pgg-cli/internal/db"
)

func main() {
	db.Connect()
	db.CreateDB()
	cmd.Execute()
}
