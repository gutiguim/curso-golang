package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println("RETORNO NOMEADO")
	soma, subtracao := calculosMatematicos(4, 5)
	println(soma, subtracao)
}
