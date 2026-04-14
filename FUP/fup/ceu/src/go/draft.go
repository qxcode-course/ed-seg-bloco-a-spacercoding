package main
import "fmt"
func main() {
    var posPedra int

    fmt.Scan(&posPedra)

    fmt.Print("[ ")
    for i := 0; i < 11; i++ {

        if i == posPedra {
            continue
        }
        
        if i == 10 {

            fmt.Print("ceu ")
            continue

        }

        fmt.Printf("%d ", i)

    }
    fmt.Println("]")
}
