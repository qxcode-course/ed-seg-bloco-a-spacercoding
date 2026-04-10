package main

import "fmt"

func quemGanhou(quempar, qtd1, qtd2 int) int {

    if (quempar == 0 && (qtd1 + qtd2) % 2 == 0) || (quempar == 1 && (qtd1 + qtd2) % 2 != 0) {return 0}

    return 1
}

func main() {
    var quemPar, qtd1, qtd2 int

    fmt.Scan(&quemPar, &qtd1, &qtd2)

    fmt.Println(quemGanhou(quemPar, qtd1, qtd2))
}
