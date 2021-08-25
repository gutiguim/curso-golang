package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("INTERFACES COMO TIPO GENÉRICO")

	generica("String")
	generica(1)
	generica(true)

	fmt.Println("Println é um exemplo pode receber muitas interfaces genéricas")
	fmt.Println(1, true, "string", float64(12345))

	mapa := map[interface{}]interface{}{
		1:            "string",
		float64(100): true,
		"String":     "string",
		true:         float64(81273),
	}

	fmt.Println(mapa)
}
