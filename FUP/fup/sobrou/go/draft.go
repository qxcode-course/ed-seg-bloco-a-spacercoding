package main
import "fmt"
func main() {
    var qtd1, qtd2, qtd3, val1, val2, val3, disponivel, total float64

    fmt.Scan(&qtd1, &qtd2, &qtd3, &val1, &val2, &val3, &disponivel)

    total = (qtd1 * val1) + (qtd2 * val2) + (qtd3 * val3)

    fmt.Printf("%.2f\n", disponivel - total)

}
