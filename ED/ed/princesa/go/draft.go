package main

import (
	"fmt"
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

func josephus(termoInicio int, jogo []jogador) {
	var varaux int

	fmt.Println(jogo)

	if qtdVivos(jogo) > 2 && jogo[termoInicio-1].num < len(jogo) {

		jogo[termoInicio].taVivo = false
		//toString(jogo, termoInicio)
		josephus(termoInicio, jogo)

	} else if qtdVivos(jogo) > 2 && jogo[termoInicio-1].num > len(jogo) {

		for i := range jogo {

			if jogo[i].taVivo {
				jogo[i].num = varaux
				jogo[i].taVivo = false
				//toString(jogo, termoInicio)
				break
			}

		}
		josephus(varaux, jogo)

	} else if qtdVivos(jogo) == 2 {

		jogo[termoInicio].taVivo = false
		//toString(jogo, termoInicio)

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

//func toString(vetor []jogador, termoInicio int) {}

func main() {

	var qtdPessoas, termoInicio int

	atribuirInt(&qtdPessoas)
	atribuirInt(&termoInicio)

	var jogo = make([]jogador, qtdPessoas)

	preencherVetor(jogo)

	josephus(termoInicio, jogo)

}
