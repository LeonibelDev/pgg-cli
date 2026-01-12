package utils

import (
	clipboard "github.com/tiagomelo/go-clipboard/clipboard"
)

func Copy(copy bool, password string) {
	if !copy {
		return
	}
	/*
		err := clipboard.Init()
		if err != nil {
			panic(err)
		}

		clipboard.Write(clipboard.FmtText, []byte(password))
	*/

	c := clipboard.New()

	if err := c.CopyText(password); err != nil {
		panic(err)
	}
}
