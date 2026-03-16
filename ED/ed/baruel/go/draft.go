package main

import "fmt"

func atribuirInteiro(inteiro *int) {

	fmt.Scan(inteiro)

}

func preencherAlbum(album map[int]int, figObtidas int) {

	var figurinha int

	for i := 0; i < figObtidas; i++ {

		fmt.Scan(&figurinha)
		fmt.Scan(album[figurinha])

		album[figurinha]++

	}

}

func analisarAlbum(album map[int]int, figTotal int) {

	for i := 1; i < figTotal; i++ {

		if album[i] > 0 {

			for j := 0; j < album[i]; j++ {

				if j < album[i]-1 {
					fmt.Printf("%d ", i)
				} else {
					fmt.Println(i)
				}

			}

		} else {

			fmt.Print(i)

		}

	}

}

func main() {

	var figTotal, figObtidas int
	atribuirInteiro(&figTotal)
	atribuirInteiro(&figObtidas)

	var album = make(map[int]int, figObtidas)
	preencherAlbum(album, figObtidas)

	analisarAlbum(album, figTotal)

}
