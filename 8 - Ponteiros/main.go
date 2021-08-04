package main

import "fmt"

func main() {
	fmt.Println("PONTEIROS")

	fmt.Println("variavel1 = 10, variavel2 = variavel1")
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println("variavel1++ -> muda só o valor da primeira variável")
	fmt.Println(variavel1, variavel2)

	// Ponteiro -> referência de memória
	fmt.Println("\nAplicação ponteiro")
	fmt.Println("variavel3 int, ponteiro *int -> printa valor 0")
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)
	// fmt.Println("variavel1++ -> muda só o valor da primeira variável")

	fmt.Println("variavel3 = 100, ponteiro = &variavel3")
	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)

	fmt.Println("varivael3, *ponteiro")
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	fmt.Println("varivael3 = 150 -> varivael3, *ponteiro")
	fmt.Println(variavel3, *ponteiro)
}
