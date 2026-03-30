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

func proxVivo(jogo []jogador, termoInicio int) int {
	var proxMorto int

	if !jogo[termoInicio].taVivo {

		for i := termoInicio; i < len(jogo); i++ {

			if jogo[i].taVivo {
				proxMorto = jogo[i].num
			} else {
				proxMorto = termoInicio
			}

		}

	}

	return proxMorto
}

func josephus(termoInicio int, jogo []jogador) {
	fmt.Println(toString(jogo, termoInicio))

	if qtdVivos(jogo) > 2 && jogo[termoInicio].num < len(jogo) {

		jogo[proxVivo(jogo, termoInicio)].taVivo = false
		josephus(proxVivo(jogo, termoInicio), jogo)

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

		if vetor[i].num == termoInicio {

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
