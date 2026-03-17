package main

import (
	"fmt"
	"slices"
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

	var posTemp = make([]posGomo, 0)

	if direcao == "L" {

		posTemp = append(posTemp, cobra[0])
		posTemp[0].x--

		posTemp = append(posTemp, cobra...)

		slices.Delete(posTemp, len(posTemp) - 1, len(posTemp))

	}
	if direcao == "R" {

		posTemp = append(posTemp, cobra[0])
		posTemp[0].x++

		posTemp = append(posTemp, cobra...)

		slices.Delete(posTemp, len(posTemp) - 1, len(posTemp))

	}
	if direcao == "U" {

		posTemp = append(posTemp, cobra[0])
		posTemp[0].y--

		posTemp = append(posTemp, cobra...)

		slices.Delete(posTemp, len(posTemp) - 1, len(posTemp))

	}
	if direcao == "D" {

		posTemp = append(posTemp, cobra[0])
		posTemp[0].y++

		posTemp = append(posTemp, cobra...)

		slices.Delete(posTemp, len(posTemp) - 1, len(posTemp))

	}

	for i := 0; i < len(posTemp) - 1; i++ {
		fmt.Printf("%d %d\n", posTemp[i].x, posTemp[i].y)
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
