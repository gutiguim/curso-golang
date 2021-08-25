package main

import "fmt"

var n int

func init() {
	fmt.Println("EXECUTANDO A FUNÇÃO INIT")
	fmt.Println("Setando n pra 10")
	n = 10
}

func main() {
	fmt.Println("EXECUTANDO A FUNÇÃO MAIN")
	fmt.Println(n)
}
