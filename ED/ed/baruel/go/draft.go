package main

import "fmt"

func atribuirInteiro(inteiro *int) {

	fmt.Scan(inteiro)

}

func preencherAlbum(album map[int]int, figObtidas, figTotal int) {

	var figurinha int

	for i := range figTotal {

		album[i] = 0

	}

	for i := 0; i < figObtidas; i++ {

		fmt.Scan(&figurinha)

		album[figurinha]++

	}

}

func analisarAlbum(album map[int]int) {

	for i := range album {

		if album[i] > 0 {

		}

	}

}

func main() {

	var figTotal, figObtidas int
	atribuirInteiro(&figTotal)
	atribuirInteiro(&figObtidas)

	var album = make(map[int]int, figObtidas)
	preencherAlbum(album, figObtidas, figTotal)

	analisarAlbum(album)

}
