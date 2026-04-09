package main
import "fmt"
func main() {
    var tCelcius, tF float64

    fmt.Scan(&tCelcius)

    tF = 1.8 * tCelcius + 32

    fmt.Printf("%.6f\n", tF)
}
