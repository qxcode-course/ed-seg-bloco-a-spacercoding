package main

import (
    "fmt"
    "math"
)

func main() {
    var l1, l2, l3, semiP, area float64

    fmt.Scan(&l1, &l2, &l3)
    
    semiP = (l1 + l2 + l3) / 2

    area = math.Sqrt(semiP * (semiP-l1) * (semiP-l2) * (semiP-l3))

    fmt.Printf("%.2f\n", area)
}
