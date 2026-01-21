package flags

import (
	"encoding/json"
	"os"

	"github.com/leonibeldev/pgg-cli/internal/db"
)

func Export(export bool) {
	if export {

		passwords := db.GetPasswords()

		// Generate .json file
		file, err := os.Create("passwords.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")
		encoder.Encode(passwords)
	}
}
