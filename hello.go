package main

import "fmt"

var a string

func main() {
	a = "String"

	if a == "String" {
		fmt.Println("é String mesmo")
	} else {
		fmt.Println("não é String")
	}
}
