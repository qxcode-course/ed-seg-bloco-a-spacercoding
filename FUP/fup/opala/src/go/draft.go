package main
import "fmt"
func main() {
    var velocidade, tempo, consumo float64

    fmt.Scan(&velocidade, &tempo, &consumo)

    horas := tempo / 60
    distancia := velocidade * horas
    desempenho := distancia / consumo

    fmt.Printf("%.2f\n", desempenho)
}
