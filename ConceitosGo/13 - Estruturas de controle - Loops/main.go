package main

import (
	"fmt"
)

func main() {
	fmt.Println("LOOPS")

	i := 0
	for i < 1 {
		fmt.Println("Incrementando I")
		i++
		// time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 3; j += 2 {
		fmt.Println("Incrementando j", j)
		// time.Sleep(time.Second)
	}

	fmt.Println("\nIterando um array")
	nomes := [3]string{"João", "Davi", "Lucas"}

	fmt.Println("indice, nome")
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("\n_, nome")
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("\nIterando uma string")

	fmt.Println("indice, string(letra)")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	fmt.Println("\nIterando um map")
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	fmt.Println("chave, valor")
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	fmt.Println("\nNão podemos usar o RANGE em uma struct")
}
