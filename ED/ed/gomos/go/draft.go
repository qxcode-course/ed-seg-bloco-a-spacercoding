package main

import (
	"fmt"
)

type posGomo struct {
	x, y int
}

func atribuirInt(meuInt *int) {

	fmt.Scan(meuInt)

}

func atribuirString(minhaRune *string) {

	fmt.Scan(minhaRune)

}

func atribuirVetor(cobra []posGomo) {

	for i := range cobra {

		fmt.Scan(&cobra[i].x, &cobra[i].y)

	}

}

func atualizarPos(cobra []posGomo, direcao string) {

	var posTemp posGomo

	if direcao == "L" {

		posTemp = cobra[0]
		cobra[0].x--

	}
	if direcao == "R" {

		cobra[0].x++

	}
	if direcao == "U" {

		cobra[0].y--

	}
	if direcao == "D" {

		cobra[0].y++

	}
	for i := range cobra {
		fmt.Printf("%d %d\n", cobra[i].x, cobra[i].y)
	}

}

func main() {

	var qtdGomos int
	atribuirInt(&qtdGomos)

	var direcao string
	atribuirString(&direcao)

	var cobra = make([]posGomo, qtdGomos)
	atribuirVetor(cobra)

	atualizarPos(cobra, direcao)

}
