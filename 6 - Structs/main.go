package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("STRUCTS")

	fmt.Println("\nDeclaração")
	var u usuario
	fmt.Println("Usuario vazio")
	fmt.Println(u)

	fmt.Println("Usuário populado usando .")
	u.nome = "Gustavo"
	u.idade = 28
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	fmt.Println("Usuário populado usando inferência de tipos")
	u2 := usuario{"Gustavo", 28, enderecoExemplo}
	fmt.Println(u2)

	fmt.Println("Populando único dado do usuário usando sintaxe u3 := usuario{nome: ''}")
	u3 := usuario{nome: "Gustavo"}
	fmt.Println(u3)
}
