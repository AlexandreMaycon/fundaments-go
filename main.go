package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Print(validateEmail())
}

func validateEmail() string {
	err := checkmail.ValidateFormat("alex123@gmail.com")

	if err != nil {
		return "O email é inválido"
	}

	return "O email é válido"
}
