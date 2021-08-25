package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("ARRAYS E SLICES")

	fmt.Println("var array1 [5]int")
	var array1 [5]int
	fmt.Println(array1)

	fmt.Println("array1[0] = 1")
	array1[0] = 1
	fmt.Println(array1)

	fmt.Println("array2 := [5]string{Posicao1, Posicao2, Posicao3, Posicao4, Posicao5}")
	array2 := [5]string{"Posicao1", "Posicao2", "Posicao3", "Posicao4", "Posicao5"}
	fmt.Println(array2)
	// array2[5] = Posicao6 -> erro array é inflexivel

	fmt.Println("\nAlocando tamanho do array de acordo com parâmetros")
	fmt.Println("array3 := [...]int{1,2,3,4,5}")
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	fmt.Println("\nCriando slices")
	slice1 := []int{10, 11, 12, 13, 14}
	fmt.Println(slice1)

	fmt.Println("Typeof slice, Typeof array3")
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	fmt.Println("Append 15 on slice1")
	slice1 = append(slice1, 15)
	fmt.Println(slice1)

	//indice inicial INCLUSIVO, indice final EXCLUSIVO (nesse caso pega pos 1 e não pega pos 3)
	fmt.Println("slice2 := array2[1:3] -> cria um ponteiro para os endereços do array2")
	slice2 := array2[1:3]
	fmt.Println(slice2)

	fmt.Println("array2[1] = Posicao alterada")
	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	fmt.Println("\n\nARRAYS INTERNOS")

	// params tipo, tamanho (quantidades de item), capacidade (quantidade máxima, opcional)
	fmt.Println("slice3 := make([]float32, 10, 11)")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("slice3 => appends para testar capacidade")
	slice3 = append(slice3, 5)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println(slice3)
	slice3 = append(slice3, 4)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println("estourando a capacidade, ele cria um outro array e dobra a capacidade")

	fmt.Println("slice4 := make([]float32, 5)")
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 4)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
