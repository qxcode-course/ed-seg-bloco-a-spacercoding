package main

import "fmt"

func quantasViagens(qtdMax, qtdTotal int) int {

    if qtdMax % qtdTotal == 0 {
        return qtdMax / qtdTotal
    }

    return (qtdMax / qtdTotal) + 1
}

func main() {
    var qtdMax, qtdTotal int

    fmt.Scan(&qtdMax, &qtdTotal)

    fmt.Println(qtdMax % qtdTotal, quantasViagens(qtdMax-1, qtdTotal))
}
