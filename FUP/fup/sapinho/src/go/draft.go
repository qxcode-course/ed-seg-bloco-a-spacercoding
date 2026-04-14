package main
import "fmt"
func main() {
    var profundidade, pulo, escorrego, pos int

    pos = 0
    fmt.Scan(&profundidade, &pulo, &escorrego)

    for {

        fmt.Printf("%d ", pos)
        if pos + pulo >= profundidade{break}
        pos += pulo
        fmt.Println(pos)
        pos -= escorrego

    }
    fmt.Println("saiu")

}
