package main
import "fmt"

func verificarSoma(somaDesejada int, subConjunto []int) bool {
    if len(subConjunto) == 1 && subConjunto[0] != somaDesejada {return false}
    somaAtual := subConjunto[0]

    if somaAtual + subConjunto[1] == somaDesejada {return true
    } else if somaAtual + subConjunto[1] > somaDesejada {

        verificarSoma(somaDesejada, subConjunto[1:])

    }

    verificarSoma(somaDesejada - somaAtual, subConjunto[1:])

    return false
}

func main() {
    var nElem, somaDesejada int
    fmt.Scan(&nElem, &somaDesejada)

    subConjunto := make([]int, nElem)

    for i := range nElem {fmt.Scan(subConjunto[i])}

    if verificarSoma(somaDesejada, subConjunto) {
        fmt.Println(true)
    } else {
        fmt.Println(false)
    }
}