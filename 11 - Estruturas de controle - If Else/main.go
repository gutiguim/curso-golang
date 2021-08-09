package main

import "fmt"

func main() {
	fmt.Println("ESTRUTURAS DE CONTROLE")

	fmt.Println("\nif numero > 15")
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	fmt.Println("\nif outro_numero := numero; outro_numero > 0 {")
	if outro_numero := numero; outro_numero > 0 {
		fmt.Println("Número é maior que zero")
	} else if outro_numero < -10 {
		fmt.Println("Número é menor que - 10")
	} else {
		fmt.Println("Número está entre -10 e 0")
	}
	fmt.Println("outro_numero é limitado ao escopo do if/else")
}
