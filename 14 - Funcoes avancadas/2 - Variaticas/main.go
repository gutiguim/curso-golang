package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println("CHEGOU UM SLICE")
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("VARIÁTICAS")

	total := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	totalVazio := soma()
	fmt.Println(totalVazio)

	fmt.Println("\nEnviando texto, mais números")
	escrever("Teste, ", 0, 1, 2, 3, 4, 5)
}
