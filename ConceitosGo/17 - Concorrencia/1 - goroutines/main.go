package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("GOROUTINES")

	go escrever("Olá mundo!") // goroutine
	escrever("Programando em go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
