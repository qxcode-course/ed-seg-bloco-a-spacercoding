package main

import "fmt"

func deletar(fila map[int]int, desistentes []int) {
	var varAux = make([]int, 0)

	for i := range fila {

		if fila[i] == 

	}

}

func main() {
	var qtdFila, qtdDesistentes int

	fmt.Scan(&qtdFila)

	var fila = make(map[int]int, qtdFila)

	for i := range qtdFila {
		var ident int

		fmt.Scan(&ident)

		fila[i] = ident
	}

	fmt.Scan(&qtdDesistentes)

	var desistentes = make([]int, qtdDesistentes)

	for i := range qtdDesistentes {
		var ident int

		fmt.Scan(&ident)

		desistentes[i] = ident
	}

	fmt.Println(deletar(fila, desistentes))
}
