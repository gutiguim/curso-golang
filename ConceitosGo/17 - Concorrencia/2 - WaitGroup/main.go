package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("WAIT GROUP")

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
	escrever("Acabou o wait group")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
