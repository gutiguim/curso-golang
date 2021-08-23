package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("MÉTODOS")

	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
}
