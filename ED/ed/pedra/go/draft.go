package main

import "fmt"

type competidor struct {
	jogada1, jogada2 int
}

func atribuirValor(numCompetidores *int) {

	fmt.Scan(numCompetidores)

}

func atribuirJogadas(jogada1, jogada2 *int) {

	fmt.Scan(jogada1, jogada2)

}

func resultadoCompeticao(classificados []competidor) int {

	var ptGanhador, ganhador int
	ptGanhador = 100
	ganhador = 100

	for i := range len(classificados) {
		if classificados[i].jogada1 >= 10 && classificados[i].jogada2 >= 10 {
			if classificados[i].jogada1 > classificados[i].jogada2 {

				if classificados[i].jogada1-classificados[i].jogada2 < ptGanhador {
					ptGanhador = classificados[i].jogada1 - classificados[i].jogada2
					ganhador = i
				}

			} else {
				ptGanhador = classificados[i].jogada2 - classificados[i].jogada1
				ganhador = i
			}
		}

	}

	return ganhador

}

func main() {

	var numCompetidores int
	atribuirValor(&numCompetidores)

	var competidores = make([]competidor, numCompetidores)
	for i := range numCompetidores {

		atribuirJogadas(&competidores[i].jogada1, &competidores[i].jogada2)

	}

	if resultadoCompeticao(competidores) < 100 {
		fmt.Println(resultadoCompeticao(competidores))
	} else {
		fmt.Println("sem ganhador")
	}

}
