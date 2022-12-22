package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main.go file")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("valid@gmail.com")

	fmt.Println(err)
}