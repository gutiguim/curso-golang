package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	// map[tipo_da_chave]tipo_dos_valores
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println("usuario['nome'], usuario.nome -> erro")
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "1",
		},
	}

	fmt.Println("usuario 2 map de map")
	fmt.Println(usuario2)

	fmt.Println("\napagar uma chave (curso) delete(usuario2, 'nome')")
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	fmt.Println("\nadicionar uma chave usuario2['signo'] = map[string]string{...")
	usuario2["signo"] = map[string]string{
		"nome": "gemeos",
	}
	fmt.Println(usuario2)
}
