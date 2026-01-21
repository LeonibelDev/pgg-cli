package flags

import (
	"fmt"

	"github.com/leonibeldev/pgg-cli/internal/db"
)

func Delete(ID int) {
	if ID != 0 {
		var response string
		fmt.Printf("Are you sure you want to delete password with ID %v? [y/N]: ", ID)
		fmt.Scanf("%s", &response)

		if response == "y" {
			db.DeletePassword(ID)
		}
	}
}
