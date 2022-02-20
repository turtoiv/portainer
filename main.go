package main

import (
	"fmt"
	"os"
	validator "validator/validator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: validator.exe containerrepr")
		return
	}

	value := validator.Validate(os.Args[1])

	fmt.Printf("value=%t", value)
}
