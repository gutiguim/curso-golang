package main

import "fmt"

func main() {
	fmt.Println("ANONIMAS")

	func() {
		fmt.Println("Anonima sem parâmetro")
	}()

	func(texto string) {
		fmt.Printf("Anonima com parâmetro: %s", texto)
	}("parâmetro\n")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("parâmetro")

	fmt.Println(retorno)
}
