package main

import (
	"fmt"
	"strconv"
	"strings"
)

func atribuirInteiro(inteiro *int) {

	fmt.Scan(inteiro)

}

func preencherAlbum(album map[int]int, figObtidas, figTotal int) {

	var figurinha int

	for i := 1; i <= figTotal; i++ {

		album[i] = 0

	}

	for i := 0; i < figObtidas; i++ {

		fmt.Scan(&figurinha)

		album[figurinha]++

	}

}

func analisarAlbum(album map[int]int, figTotal int) {

	var repetidas = make([]int, 0)
	var faltando = make([]int, 0)

	for i := 1; i < figTotal+1; i++ {

		if album[i] > 1 {
			for j := 0; j < album[i]-1; j++ {
				repetidas = append(repetidas, i)
			}

		} else if album[i] == 0 {

			faltando = append(faltando, i)

		}

	}
	if len(repetidas) > 0 {
		fmt.Println(strings.Join(formatarVetor(repetidas), " "))
	} else {
		fmt.Println("N")
	}

	if len(faltando) > 0 {
		fmt.Println(strings.Join(formatarVetor(faltando), " "))
	} else {
		fmt.Println("N")
	}

}

func formatarVetor(vetor []int) []string {

	var vetorFormatado = make([]string, 0)

	for i := range vetor {

		vetorFormatado = append(vetorFormatado, strconv.Itoa(vetor[i]))

	}

	return vetorFormatado

}

func main() {

	var figTotal, figObtidas int
	atribuirInteiro(&figTotal)
	atribuirInteiro(&figObtidas)

	var album = make(map[int]int, figObtidas)
	preencherAlbum(album, figObtidas, figTotal)

	analisarAlbum(album, figTotal)

}
