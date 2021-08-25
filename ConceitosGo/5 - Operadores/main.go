package main

import (
	"fmt"
)

func main() {
	fmt.Println("OPERADORES")
	// aritméticos +, -, *, /, %
	fmt.Println("\nAritméticos:")
	soma := 10 + 1
	fmt.Println("soma")
	fmt.Println(soma)

	subtracao := 10 - 1
	fmt.Println("subtracao")
	fmt.Println(subtracao)

	divisao := 10 / 2
	fmt.Println("divisao")
	fmt.Println(divisao)

	multiplicacao := 10 * 3
	fmt.Println("multiplicacao")
	fmt.Println(multiplicacao)

	modulo := 10 % 3
	fmt.Println("modulo")
	fmt.Println(modulo)

	// OBS: não se pode somar tipos diferentes int16 + int32 => ERRO
	fmt.Println("soma com atribuição")
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	fmt.Println("\nAtribuição")
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUICAO

	// OPERADORES RELACIONAIS
	fmt.Println("\nRelacionais")
	fmt.Println("1 > 2")
	fmt.Println(1 > 2)
	fmt.Println("1 >= 2")
	fmt.Println(1 >= 2)
	fmt.Println("1 == 2")
	fmt.Println(1 == 2)
	fmt.Println("1 <= 2")
	fmt.Println(1 <= 2)
	fmt.Println("1 < 2")
	fmt.Println(1 < 2)
	fmt.Println("1 != 2")
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("\nLógicos")
	verdadeiro, falso := true, false
	fmt.Println("verdadeiro && falso")
	fmt.Println(verdadeiro && falso)
	fmt.Println("verdadeiro || falso")
	fmt.Println(verdadeiro || falso)
	fmt.Println("!verdadeiro")
	fmt.Println(!verdadeiro)
	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	fmt.Println("\nUnários")
	fmt.Println("10++")
	numero3 := 10
	numero3++
	fmt.Println(numero3)
	fmt.Println("+=15")
	numero3 += 15
	fmt.Println(numero3)
	fmt.Println("26--")
	numero3--
	fmt.Println(numero3)
	fmt.Println("-=20")
	numero3 -= 20
	fmt.Println(numero3)
	fmt.Println("*=3")
	numero3 *= 3
	fmt.Println(numero3)
	fmt.Println("/=5")
	numero3 /= 5
	fmt.Println(numero3)
	fmt.Println("%=3")
	numero3 %= 3
	fmt.Println(numero3)
	// FIM DOS OPERADORES UNÁRIOS

	// OPERADORES TERNÁRIO
	fmt.Println("\nTernário")
	fmt.Println("NÃO EXISTE EM GO")

	fmt.Println("Necessário usar IF ELSE")
	var texto string
	if numero3 > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
	// FIM OPERATIORES TERNÁRIO
}
