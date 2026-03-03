package main

import ( 
    "fmt"
    "math"
)

func hedron(lado1, lado2, lado3 float64) (float64) {

    var semiPerimetro, area float64

    semiPerimetro = (lado1 + lado2 + lado3) / 2

    area = math.Pow(2, semiPerimetro * (semiPerimetro - lado1) * (semiPerimetro - lado2) * (semiPerimetro - lado3))

    return area

}

func main() {

    var lado1, lado2, lado3 float64

    fmt.Scan(&lado1, &lado2, &lado3)

    fmt.Println(hedron(lado1, lado2, lado3))

}
