package flags

import (
	"fmt"
	"os"
	"strings"
)

var lower string = "abcdefghijklmnopqrstuvwxyz"
var upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numbers string = "0123456789"
var special string = "!@#$%^&*()_+-=[]{}|;:,.<>/?`~"

var TypesToUse strings.Builder

func Types(types []string) (strings.Builder, error) {
	for x := range types {
		switch types[x] {
		case "lower":
			TypesToUse.WriteString(lower)
		case "upper":
			TypesToUse.WriteString(upper)
		case "numbers":
			TypesToUse.WriteString(numbers)
		case "special":
			TypesToUse.WriteString(special)
		default:
			fmt.Println("Type not found")
			os.Exit(1)
		}
	}

	return TypesToUse, nil
}
