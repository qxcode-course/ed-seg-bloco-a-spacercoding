package main
import "fmt"
func main() {
    var numImpar int

    fmt.Scan(&numImpar)

    for i := 1; i <= numImpar; i += 2 {

        fmt.Println(i)

    }

    for i := numImpar; i >= 0; i-- {

        if i % 2 == 0 {
            fmt.Println(i)
        }

    } 
}
