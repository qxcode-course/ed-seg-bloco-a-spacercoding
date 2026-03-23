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

	}

}

func josephus(termoInicio int, jogo []jogador) {

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

func toString(vetor []jogador) {

	var retorno = make([]string, 0)

	for i := range vetor {
		if vetor[i].taVivo {
			retorno = append(retorno, strconv.Itoa(vetor[i].num))
		}

	}

	fmt.Printf("[ ")
	for i := range retorno {
		fmt.Printf("%s ", retorno[i])
	}
	fmt.Println("]")

}

func main() {

	var qtdPessoas, termoInicio int

	atribuirInt(&qtdPessoas)
	atribuirInt(&termoInicio)

	var jogo = make([]jogador, qtdPessoas)

	preencherVetor(jogo)

	josephus(termoInicio, jogo)

}
