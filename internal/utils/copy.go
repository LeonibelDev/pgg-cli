package utils

import (
	"fmt"

	clipboard "github.com/tiagomelo/go-clipboard/clipboard"
)

func Copy(copy bool, password string) {
	if !copy {
		return
	}

	c := clipboard.New()

	if err := c.CopyText(password); err != nil {
		panic(err)
	}

	fmt.Println("Password copied to clipboard!")
}
