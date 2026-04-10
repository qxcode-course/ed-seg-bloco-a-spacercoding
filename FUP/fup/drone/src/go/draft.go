package main

import "fmt"

func sePassa(lado1,lado2, lado3, alturaJanela, larguraJanela int) string {

    if (lado1 * lado2 > alturaJanela * larguraJanela)  && (lado2 * lado3 > alturaJanela * larguraJanela) && (lado1 * lado3 > alturaJanela * larguraJanela) {
        return "N"
    }

    return "S"
}

func main() {
    var lado1, lado2, lado3, alturaJanela, larguraJanela int

    fmt.Scan(&lado1, &lado2, &lado3, &alturaJanela, &larguraJanela)

    fmt.Println(sePassa(lado1,lado2, lado3, alturaJanela, larguraJanela))
}
