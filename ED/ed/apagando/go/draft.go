package main

import "fmt"

func atribuirInt(endereco *int) {

	fmt.Scan(endereco)

}

func updtFila(fila, desistentes map[int]int) {
	var filaAtualizada = make([]string, 0)

	for i := range fila {

		if fila[i] == desistentes[i] {

			fila[i] = 0

		} else {

			filaAtualizada = append(filaAtualizada, fmt.Sprint(fila[i]))

		}

	}

	for i := 0; i < len(filaAtualizada)-1; i++ {

		fmt.Printf("%s ", filaAtualizada[i])

	}

	fmt.Printf("%s \n", filaAtualizada[len(filaAtualizada)-1])

}

func main() {

	var qtdPessoas, qtdDesistentes int
	var fila = make(map[int]int, qtdPessoas)

	atribuirInt(&qtdPessoas)

	for i := 0; i < qtdPessoas; i++ {
		var x int
		atribuirInt(&x)

		fila[x] = x

	}

	atribuirInt(&qtdDesistentes)

	var desistentes = make(map[int]int, qtdDesistentes)

	for i := 0; i < qtdDesistentes; i++ {
		var x int
		atribuirInt(&x)

		desistentes[x] = x

	}

	updtFila(fila, desistentes)

}
