package main

import (
	"fmt"
	"strconv"
)

type jogador struct {
	num    int
	taVivo bool
}

func atribuirInt(inteiro *int) {

	fmt.Scan(inteiro)

}

func preencherVetor(vetor []jogador) {

	for i := range vetor {

		vetor[i].num = i + 1
		vetor[i].taVivo = true

	}

}

func proxVivo(jogo []jogador, possivelVivo int) int {
	var proxMorto int = -2

	for i := possivelVivo; i < len(jogo); i++ {
		if jogo[i].taVivo {
			proxMorto = i
			return proxMorto
		}
	}
	for i := 0; i < possivelVivo; i++ {
		if jogo[i].taVivo {
			proxMorto = i
			return proxMorto
		}
	}

	return 0

}

func josephus(numInicio int, jogo []jogador) {
	fmt.Println(toString(jogo, numInicio))

	if qtdVivos(jogo) > 2 {

		jogo[proxVivo(jogo, numInicio+1)].taVivo = false
		numInicio = proxVivo(jogo, numInicio+1)
		josephus(numInicio, jogo)

	} else if qtdVivos(jogo) == 2 {

		jogo[proxVivo(jogo, numInicio+1)].taVivo = false
		numInicio = proxVivo(jogo, numInicio+1)
		fmt.Println(toString(jogo, numInicio))

	}

}

func qtdVivos(pessoas []jogador) int {
	var numVivos int

	for i := range pessoas {

		if pessoas[i].taVivo {

			numVivos++

		}

	}

	return numVivos

}

func toString(vetor []jogador, termoInicio int) []string {
	var vetorString = make([]string, 0)

	vetorString = append(vetorString, "")
	for i := range vetor {

		if vetor[i].num == termoInicio+1 && vetor[i].taVivo {

			vetorString = append(vetorString, strconv.Itoa(termoInicio+1)+">")

		} else if vetor[i].taVivo {

			vetorString = append(vetorString, strconv.Itoa(vetor[i].num)+"")
		}

	}
	vetorString = append(vetorString, "")

	return vetorString

}

func main() {

	var qtdPessoas, numInicio int

	atribuirInt(&qtdPessoas)
	atribuirInt(&numInicio)
	numInicio--

	var jogo = make([]jogador, qtdPessoas)

	preencherVetor(jogo)

	josephus(numInicio, jogo)

}
