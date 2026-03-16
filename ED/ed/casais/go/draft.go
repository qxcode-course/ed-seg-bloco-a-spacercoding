package main

import "fmt"

func atribuirValor(numAnimais *int) {

	fmt.Scan(numAnimais)

}

func preencherArca(arca map[int]int, numAnimais int) {

	var animal int

	for i := 0; i < numAnimais; i++ {
		fmt.Scan(&animal)

		fmt.Scan(arca[animal])

		arca[animal]++
	}

}

func montarCasais(arca map[int]int, numAnimais int) int {

	var numCasais int

	for i := range numAnimais {

		if arca[-i] > 0 {
			numCasais += (arca[i] + arca[-i]) / 2
		}

	}

	return numCasais
}

func main() {

	var numAnimais int
	atribuirValor(&numAnimais)

	var arca = make(map[int]int, numAnimais)
	preencherArca(arca, numAnimais)

	fmt.Println(montarCasais(arca, numAnimais))

}
