package main

import "fmt"

func atribuirInt(endereco *int) {

	fmt.Scan(endereco)

}

type pessoa struct {

	presente bool
	index int

}

func main() {
	var qtdFila, qtdDesistentes int

	atribuirInt(&qtdFila)

	var fila = make(map[int]pessoa, qtdFila)

	for i := range qtdFila {
		var x int

		atribuirInt(&x)

	}

	var desistentes = make(map[int]bool, qtdDesistentes)

}
