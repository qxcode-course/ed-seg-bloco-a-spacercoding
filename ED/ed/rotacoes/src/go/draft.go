package main

import (
	"fmt"
	"strconv"
)

func Rotacione(vet []int) []int {
    var varAux int = vet[len(vet)-1]
    
    for i := len(vet)-1; i > 0; i-- {

        vet[i] = vet[i-1]

    }
    vet[0] = varAux

    return vet
}

func VetToString(vet []int) string {
    var str string = "[ "

    for i := range len(vet) {

        str += strconv.Itoa(vet[i]) + " "

    }

    return str + "]"
}

func main() {
    var vetSize int

    fmt.Scan(&vetSize)

    var vetInt []int = make([]int, vetSize)

    for i := range vetSize {vetInt[i] = i+1}

    var numRotacoes int

    fmt.Scan(&numRotacoes)

    for range numRotacoes { vetInt = Rotacione(vetInt) }

    fmt.Println(VetToString(vetInt))
}