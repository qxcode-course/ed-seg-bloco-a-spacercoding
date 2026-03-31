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

func proxVivo(jogo []jogador, numInicio int) int {
	var proxMorto int

	if jogo[numInicio].taVivo && numInicio < len(jogo)-1 {
		proxMorto = jogo[numInicio].num
	} else if numInicio < len(jogo)-1 {
		for i := numInicio + 1; i < len(jogo); i++ {
			if jogo[i].taVivo {
				proxMorto = jogo[i].num
				break
			}
		}
	}

	return proxMorto
}

func josephus(numInicio int, jogo []jogador) {
	fmt.Println(toString(jogo, numInicio))
	if numInicio == 0 {
		numInicio++
	}

	if qtdVivos(jogo) > 2 && numInicio < len(jogo)-2 {

		jogo[numInicio].taVivo = false
		numInicio = proxVivo(jogo, numInicio)
		josephus(numInicio, jogo)

	} else if numInicio == len(jogo)-2 && qtdVivos(jogo) > 2 {
		jogo[numInicio].taVivo = false
		numInicio = proxVivo(jogo, 1)
		josephus(numInicio, jogo)
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

		if vetor[i].num == termoInicio && vetor[i].taVivo {

			vetorString = append(vetorString, strconv.Itoa(termoInicio)+">")

		} else if vetor[i].taVivo {
			vetorString = append(vetorString, strconv.Itoa(vetor[i].num)+"")
		}

	}
	vetorString = append(vetorString, "")

	return vetorString

}

func main() {

	var qtdPessoas, termoInicio int

	atribuirInt(&qtdPessoas)
	atribuirInt(&termoInicio)

	var jogo = make([]jogador, qtdPessoas)

	preencherVetor(jogo)

	josephus(termoInicio, jogo)

}
