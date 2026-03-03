package main

import (
    "fmt"
    "math"
)

func hedron(lado1, lado2, lado3 int) (float32) {

    var semiPerimetro, area float32

    semiPerimetro = (lado1 + lado2 + lado3) / 2

    area = math.Pow(2, semiPerimetro * (semiPerimetro - lado1) * (semiPerimetro - lado2) * (semiPerimetro - lado3))

    return area

}

func main() {

    var lado1, lado2, lado3 int

    fmt.Scan(&lado1, &lado2, &lado3)

    hedron(lado1, lado2, lado3)

}
