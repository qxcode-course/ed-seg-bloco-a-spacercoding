package main

import (
    "fmt"
    "math"
)

func calcRaizes(a, b, c float64) (float64, float64) {
    var delta, raiz1, raiz2 float64
    
    delta = math.Pow(b, 2) + (-4 * a * c)
    raiz1 = (((b * -1) + math.Sqrt(delta))/(2 * a))
    raiz2 = (((b * -1) - math.Sqrt(delta))/(2 * a))

    if delta >= 0 {
        return raiz1, raiz2
    }

    return -1, -1
}

func main() {
    var a, b, c, raiz1, raiz2 float64

    fmt.Scan(&a, &b, &c)

    raiz1, raiz2 = calcRaizes(a, b, c)

    if raiz1 == -1 && raiz2 == -1 {
        fmt.Println("nao ha raiz real")
    } else if raiz1 == raiz2 {
        fmt.Printf("%.2f\n", raiz1)
    } else {
        fmt.Printf("%.2f\n%.2f\n", raiz1, raiz2)
    }
    
}
