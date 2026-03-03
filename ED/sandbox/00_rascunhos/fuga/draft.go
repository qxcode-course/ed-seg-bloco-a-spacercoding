package main

import (
	"fmt"
	"internal/poll"
)

func calculoFuga(helicoptero, policial, fugitivo, direcao int) (rune) {

    var resultado rune

    if fugitivo * direcao + policial == 0 { resultado = 'N'
} else {resultado = 'S'}

    return resultado

}

func main() {

    var helicoptero, policial, fugitivo, direcao int

    fmt.Scan(&helicoptero, &fugitivo, &policial, &direcao)

    fmt.Println(calculoFuga(helicoptero, fugitivo, policial, direcao))

}
