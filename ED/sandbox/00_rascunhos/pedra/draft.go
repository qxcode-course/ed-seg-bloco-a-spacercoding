package main

import "fmt"

type concorrente struct {

    jogada1 int
    jogada2 int

}

func pontuar(competidores []int, concorrentes []concorrente, N int) ([]int) {

    for i := range N {

        if concorrentes[i].jogada1 >= 10 && concorrentes[i].jogada2 >= 10 {
            if concorrentes[i].jogada1 > concorrentes[i].jogada2 {
            competidores = append(competidores, concorrentes[i].jogada1 - concorrentes[i].jogada2)
            }else {competidores = append(competidores, concorrentes[i].jogada2 - concorrentes[i].jogada1)}
        }
    }

    return competidores
}

func calcularGanhador(competidores []int) (int) {

    pontuacaoMenor := 100
    ganhador := -1
    
    for i := range competidores {

        if competidores[i] < pontuacaoMenor {ganhador = i; pontuacaoMenor = competidores[i]}

    }

    return ganhador + 1

}

func main() {

    var N int
    fmt.Scan(&N)

    var concorrentes = make([]concorrente, N)
    var competidores = make([]int, 0)


    for i := range N {
        
        fmt.Scan(&concorrentes[i].jogada1, &concorrentes[i].jogada2)

    }

    

    if calcularGanhador(pontuar(competidores, concorrentes, N)) == -1 {fmt.Println("sem ganhador")
    }else {fmt.Println(calcularGanhador(pontuar(competidores, concorrentes, N)))}

}
