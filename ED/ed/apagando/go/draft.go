package main

import (
	"fmt"
)

func compararEDeletar(fila []int, desistentes map[int]int) []int {
	var vetResultante = make([]int, 0)

	for i := range fila {
		var x int

		x = fila[i]

		if desistentes[x] > 0 {

			continue

		} else {
			vetResultante = append(vetResultante, x)
		}

	}

	return vetResultante
}

func main() {
	var qtdInFila, qtdOutFila int

	fmt.Scan(&qtdInFila)

	var fila = make([]int, qtdInFila)

	for i := range qtdInFila {
		var x int

		fmt.Scan(&x)

		fila[i] = x

	}

	fmt.Scan(&qtdOutFila)

	var desistentes = make(map[int]int, qtdOutFila)

	for range qtdOutFila {
		var x int

		fmt.Scan(&x)

		desistentes[x]++

	}

	varAux := compararEDeletar(fila, desistentes)

	for i := range len(varAux) {
		fmt.Printf("%d ", varAux[i])
	}
	fmt.Println()

}
