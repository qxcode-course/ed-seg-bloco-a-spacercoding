package main

import "fmt"

func calcParteArea(inicioCorte, fimCorte int) int {
    var areaTrapezio int = ((inicioCorte + fimCorte) / 2) * 70

    if areaTrapezio > (160 * 70) / 2 {return 1
    } else if areaTrapezio < (160 * 70) / 2 {return 2}

    return 0
}

func main() {
    var inicioCorte, fimCorte int

    fmt.Scan(&inicioCorte, &fimCorte)

    fmt.Println(calcParteArea(inicioCorte, fimCorte))
}
