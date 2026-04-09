package main
import "fmt"
func main() {
    var hora, minuto, dia, mes, ano int

    fmt.Scan(&hora, &minuto, &dia, &mes, &ano)

    fmt.Printf("%.2d:%.2d %.2d/%.2d/%.2d\n", hora, minuto, dia, mes, ano % 100)
}
