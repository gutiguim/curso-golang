package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("MÉTODOS")

	fmt.Println("\nCriando métodos:")
	usuario1 := usuario{"Golias", 15}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	fmt.Println(usuario2)
	usuario2.salvar()

	fmt.Println("\nCheckando maior idade:")
	fmt.Printf("%s é maior de idade? %t\n", usuario1.nome, usuario1.maiorDeIdade())
	fmt.Printf("%s é maior de idade? %t\n", usuario2.nome, usuario2.maiorDeIdade())

	fmt.Println("\nMudando a idade (ponteiro):")
	fmt.Println(usuario1.idade)
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
