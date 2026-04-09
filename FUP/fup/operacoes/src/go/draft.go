package main
import "fmt"
func main() {
    var n1, n2 float64

    fmt.Scan(&n1, &n2)

    fmt.Printf("%.0f\n%.0f\n%.0f\n%.2f\n%d\n", n1 + n2, n1 - n2, n1 * n2, n1 / n2, int(n1) % int(n2))
}
