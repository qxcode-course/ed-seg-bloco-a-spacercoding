package main

import "fmt"

func atribuirValor(numAnimais *int) {

	fmt.Scan(numAnimais)

}

func preencherArca(arca map[int]int, numAnimais int) {

	var animal int

	for i := range numAnimais {
		fmt.Scan(&animal, arca[animal])

		if i == animal || i == -animal {
			arca[animal] += 1
		}

	}

}

func montarCasais(arca map[int]int, numAnimais int) int {

	var numCasais int

	for i := range numAnimais / 2 {
		numCasais += (arca[i] + arca[-i]) / 2
	}

	return numCasais
}

func main() {

	var numAnimais int
	atribuirValor(&numAnimais)

	var arca map[int]int
	preencherArca(arca, numAnimais)

	fmt.Println(montarCasais(arca, numAnimais))

}
