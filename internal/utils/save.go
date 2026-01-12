package utils

import "fmt"

func Save(social string, password string) {
	if social != "" {
		fmt.Printf("Social: %v\nPassword: %v\n", social, password)
	}
}
