package main
import "fmt"

func verificarSoma(somaDesejada int, subConjunto []int) bool {
    if len(subConjunto) == 1 {return false}
    var aux int = subConjunto[0]

    for i := 1; i < len(subConjunto); i++ {

        if aux + subConjunto[i] == somaDesejada {return true
        }

        

    }

    return verificarSoma(somaDesejada, subConjunto[1:])

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