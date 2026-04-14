package main
import "fmt"
func main() {
    var posPedra int
    var peInicial, outroPe string

    fmt.Scan(&posPedra, &peInicial)

    if peInicial == "e" {outroPe = "d"
    } else {outroPe = "e"}

    fmt.Print("[ ")
    for i := 0; i < 11; i++ {

        if i == posPedra {
            
            varAux := peInicial
            peInicial = outroPe
            outroPe = varAux
            continue

        } else if i == 10 {

            fmt.Print("ceu ")
            continue

        }

        if i % 2 == 0 {

            fmt.Printf("%d%s ", i, peInicial)

        } else {

            fmt.Printf("%d%s ", i, outroPe)

        }

    }
    fmt.Println("]")
}
