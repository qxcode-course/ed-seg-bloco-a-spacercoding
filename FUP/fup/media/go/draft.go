package main
import "fmt"
func main() {
    var num1, num2 float64

    fmt.Scan(&num1, &num2)

    fmt.Printf("%.1f\n", (num1 + num2) / 2)

}
