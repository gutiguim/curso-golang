package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("parâmetro função f")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// não quer o segundo valor, underline ignora resultado
	resultadoSoma1, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma1)

	// não quer o primeiro valor, underline ignora resultado
	_, resultadoSubtracao1 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao1)
}
