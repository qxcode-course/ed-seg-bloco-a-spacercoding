package main
import "fmt"
func main() {
    var tempoSeg, hora, minuto, segundo int

    fmt.Scan(&tempoSeg)
    hora = tempoSeg / 3600
    minuto = (tempoSeg % 3600) / 60
    segundo = tempoSeg % 60

    fmt.Printf("%d:%d:%d\n", hora, minuto, segundo)

}
