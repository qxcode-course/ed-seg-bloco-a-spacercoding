package main

import "fmt"

func contagem(figurinhas []int) ([]int, []int) {
    var repetidas, faltando = make([]int, 0), make([]int, 0)

    for i := range figurinhas {

        if figurinhas[i] == 0 {
            faltando = append(faltando, i+1)
        } else if figurinhas[i] > 1 {

            repetidas = append(repetidas, i+1)

        }

    }

    

    return repetidas, faltando
}

func main() {
    var total, possui int

    fmt.Scan(&total, &possui)

    var figurinhas = make([]int, total)

    for i := 0; i < total; i++ {

        figurinhas[i] = 0

    }
    
    for range possui {
        var figurinha int
        
        fmt.Scan(&figurinha)
        
        figurinhas[figurinha-1]++
    }

    vetAux, vetAux2 := contagem(figurinhas)

    if len(vetAux) == 0 {fmt.Println("N")
    } else {
    for i := range len(vetAux) - 1 {
    fmt.Printf("%d ", vetAux[i])
    }
    fmt.Println(vetAux[len(vetAux)-1])
    }

    if len(vetAux2) == 0 {fmt.Println("N")
    } else {
    for i := range len(vetAux2) - 1 {
        fmt.Printf("%d ", vetAux2[i])
    }
    fmt.Println(vetAux2[len(vetAux2)-1])
    }

}