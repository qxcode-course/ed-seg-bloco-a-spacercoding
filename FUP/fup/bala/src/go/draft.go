package main

import (
    "fmt"
    "math"
)

func main() {
    var x1, y1, x2, y2, distancia float64

    fmt.Scan(&x1, &y1, &x2, &y2)

    distancia = math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1 - y2, 2))

    fmt.Printf("%.2f\n", distancia)
}