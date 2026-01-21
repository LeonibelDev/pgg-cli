package flags

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/leonibeldev/pgg-cli/internal/db"
)

func List(list bool) {
	if list {

		passwords := db.GetPasswords()

		t := table.NewWriter()
		t.AppendHeader(table.Row{"#", "Service", "Username", "Password"})

		for x := range passwords {
			t.AppendRow(table.Row{
				passwords[x].ID,
				passwords[x].Service,
				passwords[x].Username,
				passwords[x].Password,
			})
		}

		fmt.Println(t.Render())

	}
}
