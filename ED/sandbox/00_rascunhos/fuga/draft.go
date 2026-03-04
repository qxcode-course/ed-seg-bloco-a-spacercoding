package main

import (
	"fmt"
)

func calculoFuga(helicoptero, policial, fugitivo, direcao int) (rune) {

    var resultado rune

    if direcao == 1 {

        if fugitivo > policial && fugitivo < helicoptero {resultado = 'S'
    }   else {resultado = 'N'}

    } else {

        if (fugitivo < policial && fugitivo > helicoptero) || (fugitivo < policial && helicoptero > policial) {resultado = 'S'
    }   else {resultado = 'N'}

    }

    return resultado

}

func main() {

    var helicoptero, policial, fugitivo, direcao int

    fmt.Scan(&helicoptero, &fugitivo, &policial, &direcao)

    fmt.Printf("%c\n", calculoFuga(helicoptero, fugitivo, policial, direcao))

}
