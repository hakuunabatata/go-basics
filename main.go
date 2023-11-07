package main

import (
	"fmt"
)

func main() {
	a := "String"
	v := `
		Isso aqui tem mais 
		de uma linha
	`

	if a == "String" {
		fmt.Println(math.some(2, 5))
	} else {
		fmt.Printf("%v ooo", v)
	}
}
