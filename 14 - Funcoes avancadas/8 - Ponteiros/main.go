package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("PONTEIROS")

	numero := 20
	numeroInvertido := inverterSinal(numero)

	fmt.Println("numero continua o mesmo por não utilizarmos ponteiros")
	fmt.Println("numeroInvertido")
	fmt.Println(numeroInvertido)
	fmt.Println("numero")
	fmt.Println(numero)

	fmt.Println("\nNovo número")
	novoNumero := 40
	fmt.Println(novoNumero)
	fmt.Println("novoNumero inverte por utilizarmos ponteiros")
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
