package main

import (
	"errors"
	"fmt"
)

func main() {
	// INT e UINT
	fmt.Println("\nINT e UINT:")
	// int8, 16, 32, 64
	// int -> usa a baso da arquitetura do computador
	var numero int64 = 1000000000000000

	fmt.Println(numero)

	var numero2 uint32 = 12313
	fmt.Println(numero2)

	// alias
	fmt.Println("\nALIAS -> INT e UINT:")
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	// FLOATS
	fmt.Println("\nFLOATS:")
	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123123123.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1345.67
	fmt.Println(numeroReal3)

	// STRING
	fmt.Println("\nSTRING:")
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// não existe char em go, esse é o mais próximo de char
	char := 'B'
	// printa o valor de B na tabela ASCII
	fmt.Println(char) //66

	// BOOLEANO
	fmt.Println("\nBOOLEANO:")
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	// ERROR
	fmt.Println("\nERROR:")

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	// VALOR ZERO
	fmt.Println("\nVALOR ZERO:")

	var texto string
	fmt.Println(texto)

	var inteiro16 int16
	fmt.Println(inteiro16) //0

	var fluntuantet32 float32
	fmt.Println(fluntuantet32) //0

	var booleano3 bool
	fmt.Println(booleano3) //false

	var erro2 error
	fmt.Println(erro2) //<nil>
}
